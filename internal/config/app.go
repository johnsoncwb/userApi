package config

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	AppName string

	NewRelicLiscence    string
	NewRelicApp         *newrelic.Application
	NewRelicTransaction *newrelic.Transaction
	NewRelicContext     context.Context
	NewRelicLogger      *logrus.Entry

	UsersUrl string
}

var GlobalConfig AppConfig

func (cfg *AppConfig) LoadVariables() (err error) {
	if err = godotenv.Load(); err != nil {
		return
	}
	cfg.AppName = os.Getenv("APP_NAME")
	cfg.NewRelicLiscence = os.Getenv("NEW_RELIC_LICENSE_KEY")
	cfg.UsersUrl = os.Getenv("EXTERNAL_URL")

	return
}
