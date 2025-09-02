package domain

import "time"

const (
	UserStatusPending int16 = iota
	UserStatusActive
	UserStatusDeactivate
)

type User struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Status    int16     `json:"status"`
	Password  string    `json:"password"`
	Pin       string    `json:"pin"`
	RoleID    int64     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ActivationToken     string    `json:"token_activation,omitempty"`
	ActivationExpiredAt time.Time `json:"token_activation_expired_at,omitempty"`
}

type UserSimplified struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Status    int16     `json:"status"`
	RoleID    int64     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRolePermission struct {
	ID         int64                  `json:"id"`
	FullName   string                 `json:"full_name"`
	Email      string                 `json:"email"`
	RoleName   string                 `json:"role_name,omitempty"`
	RoleID     int64                  `json:"role_id,omitempty"`
	Permission []PermissionSimplified `json:"permission,omitempty"`
}
