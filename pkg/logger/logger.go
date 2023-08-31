package logger

import (
	"go.uber.org/fx"
	uberZap "go.uber.org/zap"

	"github.com/abdivasiyev/microservice_template/pkg/logger/zap"
)

var FxOption = fx.Provide(
	fx.Annotate(zap.New, fx.As(new(Logger))),
)

type Logger interface {
	Debug(message string, fields ...uberZap.Field)
	Info(message string, fields ...uberZap.Field)
	Warn(message string, fields ...uberZap.Field)
	Error(message string, fields ...uberZap.Field)
	Fatal(message string, fields ...uberZap.Field)
	Sync() error
}
