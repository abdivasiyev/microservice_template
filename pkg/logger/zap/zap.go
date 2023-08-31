package zap

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/abdivasiyev/microservice_template/pkg/config"
	"github.com/abdivasiyev/microservice_template/pkg/env"
	"github.com/abdivasiyev/microservice_template/pkg/fs"
)

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    config.Config
	FS        fs.FS
}

type Logger struct {
	zap *zap.SugaredLogger
}

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}

// New ...
func New(params Params) (*Logger, error) {
	const (
		prodEnv = "production"
		testEnv = "test"
	)

	var (
		logLevel    = params.Config.GetString(env.LogLevel)
		namespace   = params.Config.GetString(env.AppNamespace)
		environment = params.Config.GetString(env.AppEnvironment)
		encoderCfg  = func(environment, timeFormat string) zapcore.EncoderConfig {
			var cfg zapcore.EncoderConfig

			if environment == prodEnv {
				cfg = zap.NewProductionEncoderConfig()
			} else {
				cfg = zap.NewDevelopmentEncoderConfig()
			}

			cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(timeFormat))
			}
			return cfg
		}(environment, time.DateTime)
		logFile = params.FS.RootDir() + "/logs/" + namespace + ".log"
		ll      = lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    1024, // MB
			MaxBackups: 30,
			MaxAge:     90, // days
			Compress:   true,
		}
	)

	globalLevel, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		return nil, errors.Join(err)
	}

	_ = zap.RegisterSink("lumberjack", func(*url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &ll,
		}, nil
	})

	loggerConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(globalLevel),
		Development:       environment != prodEnv,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{fmt.Sprintf("lumberjack:%s", logFile), "stderr"},
		ErrorOutputPaths:  []string{fmt.Sprintf("lumberjack:%s", logFile), "stderr"},
		DisableStacktrace: true,
	}

	if environment == testEnv {
		loggerConfig.OutputPaths = []string{}
		loggerConfig.ErrorOutputPaths = []string{}
	}

	zapLog, err := loggerConfig.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return nil, errors.Join(err)
	}

	zap.ReplaceGlobals(zapLog)()

	result := &Logger{
		zap: zapLog.Named(namespace).Sugar(),
	}

	params.Lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return result.Sync()
		},
	})

	return result, nil
}

func (l *Logger) Debug(message string, fields ...zap.Field) {
	l.zap.Desugar().Debug(message, fields...)
}

func (l *Logger) Info(message string, fields ...zap.Field) {
	l.zap.Desugar().Info(message, fields...)
}

func (l *Logger) Warn(message string, fields ...zap.Field) {
	l.zap.Desugar().Warn(message, fields...)
}

func (l *Logger) Error(message string, fields ...zap.Field) {
	l.zap.Desugar().Error(message, fields...)
}

func (l *Logger) Fatal(message string, fields ...zap.Field) {
	l.zap.Desugar().Fatal(message, fields...)
}

func (l *Logger) Sync() error {
	if err := l.zap.Desugar().Sync(); err != nil {
		return err
	}

	return l.zap.Sync()
}
