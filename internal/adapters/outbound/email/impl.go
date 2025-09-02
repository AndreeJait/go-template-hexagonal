package email

import (
	"context"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"

	"github.com/AndreeJait/go-utility/emailw"
	"github.com/AndreeJait/go-utility/loggerw"
	"net/url"
)

type impl struct {
	emailW emailw.EmailW
	logger loggerw.Logger
	cfg    *config.Config
}

func (i impl) SendEmailActivation(ctx context.Context, param SendEmailActivationParam) error {

	u, err := url.Parse(i.cfg.Service.RedirectFrontend)
	if err != nil {
		panic(err)
	}
	// Query params
	q := u.Query()
	q.Set("email", param.Email)
	q.Set("token", param.Token)
	// Encode back to RawQuery
	u.RawQuery = q.Encode()

	err = i.emailW.SentEmail(emailw.SentEmailParam{
		Subject:  "Activation account",
		Sender:   i.cfg.Service.EmailSender,
		To:       []string{param.Email},
		Template: "files/emails/activation-email.html",
		Param: map[string]interface{}{
			"name": param.Name,
			"url":  u.String(),
		},
	})
	if err != nil {
		i.logger.Errorf(ctx, err, "failed to send email")
		return err
	}
	return nil
}

func NewEmailW(emailW emailw.EmailW, log loggerw.Logger, cfg *config.Config) Email {
	return &impl{emailW: emailW, logger: log, cfg: cfg}
}
