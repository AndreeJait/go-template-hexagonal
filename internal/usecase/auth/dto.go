package auth

import (
	"fmt"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type GetUserByEmailRequest struct {
	Email string `json:"email"`
}

type CreateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	RoleID   int64  `json:"role_id"`
}

func (c CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.FullName, validation.Required),
		validation.Field(&c.RoleID, validation.Required),
	)
}

type CreateUserPasswordRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token           string `json:"token"`
	Email           string `json:"email"`
}

type CreateUserPasswordResponse struct {
	Token string `json:"token"`
}

func (c CreateUserPasswordRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Password,
			validation.Required,
		),
		validation.Field(&c.ConfirmPassword,
			validation.Required,
			validation.By(func(value interface{}) error {
				confirm := value.(string)
				if confirm != c.Password {
					return fmt.Errorf("password and confirm password do not match")
				}
				return nil
			}),
		),
		validation.Field(&c.Token, validation.Required),
		validation.Field(&c.Email, validation.Required, is.Email),
	)
}

type CreateUserPinRequest struct {
	Pin        string `json:"pin"`
	ConfirmPin string `json:"confirm_pin"`
	Token      string `json:"token"`
	Email      string `json:"email"`
}

func (c CreateUserPinRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Pin,
			validation.Required,
			is.UTFNumeric,
			validation.Length(6, 6),
		),
		validation.Field(&c.ConfirmPin,
			validation.Required,
			validation.Length(6, 6),
			is.UTFNumeric,
			validation.By(func(value interface{}) error {
				confirm := value.(string)
				if confirm != c.Pin {
					return fmt.Errorf("pin and confirm pin do not match")
				}
				return nil
			}),
		),
		validation.Field(&c.Token, validation.Required),
		validation.Field(&c.Email, validation.Required, is.Email),
	)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c LoginRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.Password, validation.Required))
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
	User      domain.UserSimplified
}
