package constant

import (
	"github.com/AndreeJait/go-utility/errow"
	"github.com/AndreeJait/go-utility/response"
)

var (
	// 404
	ErrUserNotFound = errow.ErrorW{Code: 404001, Message: "User Not Found"}

	// 400
	ErrWrongPassword     = errow.ErrorW{Code: 400001, Message: "Wrong Password"}
	ErrUserAlreadyExists = errow.ErrorW{Code: 400002, Message: "User Already Exists"}
	// 401
	ErrUserDidNotHaveAccessToThisFeature = errow.ErrorW{Code: 401001, Message: "User Did not have access to this feature"}
	ErrUserIsDeactivated                 = errow.ErrorW{Code: 401002, Message: "User is deactivated"}
)

var ErrorMap = map[errow.ErrorWCode]response.ErrResponseFunc{
	errow.ErrInternalServer.Code: response.ErrInternalServerError,

	errow.ErrSessionExpired.Code: response.ErrSessionExpired,

	errow.ErrForbidden.Code: response.ErrForbidden,

	errow.ErrUnauthorized.Code: response.ErrUnauthorized,

	errow.ErrBadRequest.Code:  response.ErrBadRequest,
	ErrWrongPassword.Code:     response.ErrBadRequest,
	ErrUserAlreadyExists.Code: response.ErrBadRequest,

	errow.ErrResourceNotFound.Code: response.ErrNotFound,
	ErrUserNotFound.Code:           response.ErrNotFound,

	errow.ErrInvalidSigningMethod.Code: response.ErrForbidden,
	errow.ErrInvalidToken.Code:         response.ErrForbidden,
	ErrUserIsDeactivated.Code:          response.ErrBadRequest,
}
