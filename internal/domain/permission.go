package domain

import "time"

type UserPermissionsName string

const (
	UserCanCreateCustomer   UserPermissionsName = "customer/create"
	UserCanUpdateCustomer   UserPermissionsName = "customer/update"
	UserCanDeleteCustomer   UserPermissionsName = "customer/delete"
	UserCanReadCustomer     UserPermissionsName = "customer/read"
	UserCanApprovalCustomer UserPermissionsName = "customer/approval"
)

type Permission struct {
	ID        int64               `json:"id"`
	Name      UserPermissionsName `json:"name"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

type PermissionSimplified struct {
	ID   int64               `json:"id"`
	Name UserPermissionsName `json:"name"`
}
