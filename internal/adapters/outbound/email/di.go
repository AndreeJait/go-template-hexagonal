package email

import "context"

type Email interface {
	SendEmailActivation(ctx context.Context, param SendEmailActivationParam) error
}
