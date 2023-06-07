package mmHttpClient

import (
	"context"
	"crypto/tls"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
	"time"
)

type HttpClient struct {
	http.Client
	ctx context.Context
}

// NewHttpClient creates an http client. This custom HttpClient has some utilities as NewRelic external segment and transport skip verify.
//
// c := mmhttpclient.NewHttpClient(ctx, 8 * time.Second).WithSkipVerify()
//
// res, err := c.Do(request)
func NewHttpClient(ctx context.Context, timeout time.Duration) *HttpClient {
	client := http.Client{
		Timeout: timeout,
	}

	return &HttpClient{
		Client: client,
		ctx:    ctx,
	}
}

// WithSkipVerify sets the default with http.DefaultTransport and sets InsecureSkipVerify to true, so the request can run as expected inside the Docker container
func (h *HttpClient) WithSkipVerify() *HttpClient {
	transport := http.DefaultTransport
	tlsClientConfig := &tls.Config{InsecureSkipVerify: true}
	transport.(*http.Transport).TLSClientConfig = tlsClientConfig
	transport.(*http.Transport).DisableKeepAlives = true

	h.Transport = transport
	return h
}

// Do is a proxy to http.Request Do method. It abstracts NewRelic ExternalSegment call if there's a transaction inside the context
func (h *HttpClient) Do(req *http.Request) (*http.Response, error) {
	txn := newrelic.FromContext(h.ctx)

	var exSeg *newrelic.ExternalSegment
	if txn != nil {
		exSeg = newrelic.StartExternalSegment(txn, req)
		defer exSeg.End()
	}

	res, err := h.Client.Do(req)

	if exSeg != nil {
		exSeg.Response = res
	}

	return res, err
}
