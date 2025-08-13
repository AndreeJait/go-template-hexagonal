package constant

import (
	"github.com/AndreeJait/go-utility/errow"
	"github.com/AndreeJait/go-utility/response"
)

var (
	ErrUserNotFound = errow.ErrorW{Code: 404001, Message: "User Not Found"}
)

var ErrorMap = map[errow.ErrorWCode]response.ErrResponseFunc{
	errow.ErrInternalServer.Code: response.ErrInternalServerError,

	errow.ErrSessionExpired.Code: response.ErrSessionExpired,

	errow.ErrForbidden.Code: response.ErrForbidden,

	errow.ErrUnauthorized.Code: response.ErrUnauthorized,

	errow.ErrBadRequest.Code: response.ErrBadRequest,

	errow.ErrResourceNotFound.Code: response.ErrNotFound,
	ErrUserNotFound.Code:           response.ErrNotFound,
}
