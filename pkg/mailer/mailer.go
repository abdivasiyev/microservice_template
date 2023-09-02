package mailer

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/microservice_template/pkg/mailer/gmail"
	"github.com/abdivasiyev/microservice_template/pkg/mailer/mail"
)

var FxOption = fx.Provide(
	fx.Annotate(gmail.New, fx.As(new(Mailer))),
)

type Mailer interface {
	Send(mail mail.Mail) error
}
