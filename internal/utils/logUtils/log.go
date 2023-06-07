package logutils

import (
	"context"

	log "github.com/sirupsen/logrus"
)

var loggerKey = "logrus_ctx"

// LoggerFromContext Return the logrus instance used on the context
func LoggerFromContext(ctx context.Context) *log.Entry {
	entry, ok := ctx.Value(loggerKey).(*log.Entry)
	if !ok {
		entry = log.WithContext(ctx)
	}

	return entry
}

// SetContextLogger Set the logrus instance to be used on the context
func SetContextLogger(ctx context.Context, logger *log.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
