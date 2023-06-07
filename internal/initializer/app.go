package initializer

import (
	"context"
	"time"

	"github.com/johnsoncwb/userApi/internal/config"
	logutils "github.com/johnsoncwb/userApi/internal/utils/logUtils"
	"github.com/newrelic/go-agent/v3/integrations/logcontext/nrlogrusplugin"
	"github.com/newrelic/go-agent/v3/newrelic"
	log "github.com/sirupsen/logrus"
)

func LoadEnv() {
	err := config.GlobalConfig.LoadVariables()
	if err != nil {
		panic(err)
	}
}

func StartApp() {
	LoadEnv()
	log.SetFormatter(nrlogrusplugin.ContextFormatter{})

	ctx := context.Background()
	ctx = startNewRelic(ctx)
	startLogger(ctx)

}

func startLogger(ctx context.Context) context.Context {

	config.GlobalConfig.NewRelicLogger = log.WithContext(ctx)
	return logutils.SetContextLogger(ctx, config.GlobalConfig.NewRelicLogger)
}

func FinishApp() {
	config.GlobalConfig.NewRelicTransaction.End()
}

func startNewRelic(ctx context.Context) context.Context {
	var err error

	config.GlobalConfig.NewRelicApp, err = newrelic.NewApplication(
		newrelic.ConfigAppName(config.GlobalConfig.AppName),
		newrelic.ConfigLicense(config.GlobalConfig.NewRelicLiscence),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		log.Println("error to initialize to newrelic")
	}

	err = config.GlobalConfig.NewRelicApp.WaitForConnection(5 * time.Second)

	if err != nil {
		log.Println("error to connect to newrelic")
	}

	txn := config.GlobalConfig.NewRelicApp.StartTransaction("App")
	ctx = newrelic.NewContext(ctx, txn)

	config.GlobalConfig.NewRelicTransaction = txn
	config.GlobalConfig.NewRelicContext = ctx

	log.Println("connected to newrelic")

	return ctx

}
