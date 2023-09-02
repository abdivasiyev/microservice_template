package gmail

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"

	"go.uber.org/fx"

	"github.com/abdivasiyev/microservice_template/pkg/config"
	"github.com/abdivasiyev/microservice_template/pkg/env"
	"github.com/abdivasiyev/microservice_template/pkg/mailer/mail"
)

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    config.Config
}

type Gmail struct {
	config config.Config
	client *smtp.Client
}

func New(params Params) (*Gmail, error) {
	g := &Gmail{
		config: params.Config,
	}

	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			client, err := g.getClient()
			if err != nil {
				return errors.Join(err)
			}

			g.client = client
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return g.client.Close()
		},
	})

	return g, nil
}

func (g *Gmail) getClient() (*smtp.Client, error) {
	var (
		smtpAddr  = fmt.Sprintf("%s:%d", g.config.GetString(env.SmtpHost), g.config.GetInt(env.SmtpPort))
		tlsConfig = &tls.Config{ServerName: g.config.GetString(env.SmtpHost), InsecureSkipVerify: true}
	)

	conn, err := tls.Dial("tcp", smtpAddr, tlsConfig)
	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, g.config.GetString(env.SmtpHost))
	if err != nil {
		return nil, err
	}

	var (
		auth = smtp.PlainAuth(
			"",
			g.config.GetString(env.SmtpUsername),
			g.config.GetString(env.SmtpPassword),
			g.config.GetString(env.SmtpHost),
		)
	)

	if err = g.client.Auth(auth); err != nil {
		return nil, errors.Join(err)
	}

	return client, nil
}

func (g *Gmail) Send(mail mail.Mail) error {
	if err := g.client.Mail(g.config.GetString(env.SmtpUsername)); err != nil {
		return errors.Join(err)
	}

	for _, toEmail := range mail.To {
		if err := g.client.Rcpt(toEmail); err != nil {
			return errors.Join(err)
		}
	}

	writer, err := g.client.Data()
	if err != nil {
		return errors.Join(err)
	}
	defer writer.Close()

	n, err := writer.Write(mail.Build(g.config.GetString(env.SmtpUsername)))
	if err != nil {
		return errors.Join(err)
	}

	if n == 0 {
		return errors.Join(errors.New("gmail: 0 byte written to client"))
	}

	return nil
}
