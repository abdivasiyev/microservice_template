package sentry

import (
	"errors"
	"fmt"

	"github.com/getsentry/sentry-go"
	"go.uber.org/fx"

	"github.com/abdivasiyev/microservice_template/pkg/config"
	"github.com/abdivasiyev/microservice_template/pkg/env"
)

type Params struct {
	fx.In
	Config config.Config
}

type Sentry struct {
	hub         *sentry.Hub
	environment string
}

func New(params Params) (*Sentry, error) {
	debug := false

	if params.Config.GetString(env.AppEnvironment) == "development" {
		debug = true
	}

	sentryClient, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:              params.Config.GetString(env.SentryDSN),
		Environment:      params.Config.GetString(env.AppEnvironment),
		Debug:            debug,
		AttachStacktrace: true,
		TracesSampleRate: 1,
		ServerName:       params.Config.GetString(env.AppNamespace),
	})
	if err != nil {
		return &Sentry{}, errors.Join(err)
	}

	hub := sentry.CurrentHub()
	hub.BindClient(sentryClient)

	return &Sentry{
		hub:         hub,
		environment: params.Config.GetString(env.AppEnvironment),
	}, nil
}

func (s *Sentry) isProd() bool {
	return s.environment == "production"
}

func (s *Sentry) Message(format string, args ...any) {
	if !s.isProd() {
		return
	}
	s.hub.CaptureMessage(fmt.Sprintf(format, args...))
}

func (s *Sentry) Error(err error) {
	if !s.isProd() {
		return
	}
	s.hub.CaptureException(err)
}
