package email

type SendEmailActivationParam struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
