package main

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/microservice_template/pkg/config"
	"github.com/abdivasiyev/microservice_template/pkg/fs"
	"github.com/abdivasiyev/microservice_template/pkg/logger"
	"github.com/abdivasiyev/microservice_template/pkg/mailer"
	"github.com/abdivasiyev/microservice_template/pkg/monitoring"
)

var Options = []fx.Option{
	config.FxOption,
	fs.FxOption,
	logger.FxOption,
	monitoring.FxOption,
	mailer.FxOption,
}

func main() {
	app := fx.New(Options...)
	app.Run()
}
