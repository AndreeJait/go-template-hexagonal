package auth

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http/common/middleware"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/AndreeJait/go-template-hexagonal/internal/usecase"
	"github.com/AndreeJait/go-template-hexagonal/internal/usecase/auth"
	"github.com/AndreeJait/go-utility/response"
	"github.com/AndreeJait/go-utility/tracer"
	"github.com/labstack/echo/v4"
)

type handler struct {
	route *echo.Group
	uc    *usecase.UseCase
	cfg   *config.Config
}

func NewAuthHandler(cfg *config.Config, route *echo.Group, uc *usecase.UseCase) http.Handler {
	return &handler{
		route: route,
		uc:    uc,
		cfg:   cfg,
	}
}

func (h *handler) Handle() {
	groupPublic := h.route.Group("/auth")
	groupInternal := h.route.Group("/internal/auth")

	groupPublic.POST("/create-password", h.createPassword)
	groupPublic.POST("/create-pin", h.createPin)
	groupPublic.POST("/login", h.login)

	groupInternal.Use(middleware.BasicAuthLogged(h.cfg))
	groupInternal.POST("/account", h.internalCreateNewAccount)

}

// internalCreateNewAccount godoc
// @Summary      Create new user from internal
// @Description  Create a new user account from internal
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Param        request  body  auth.CreateUserRequest  true  "CreateUserRequest"
// @Success      200  {object}  response.Response{}  "success create user"
// @Failure      400  {object}  response.ErrorResponse      "validation/bind error"
// @Failure      409  {object}  response.ErrorResponse      "email already exists"
// @Failure      500  {object}  response.ErrorResponse      "internal error"
// @Router       /internal/auth/account [post]  // <-- adjust to your actual route
func (h *handler) internalCreateNewAccount(c echo.Context) error {
	ctx := c.Request().Context()
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(h.internalCreateNewAccount))
	defer span.End()

	param := auth.CreateUserRequest{}
	if err := c.Bind(&param); err != nil {
		return err
	}

	if err := param.Validate(); err != nil {
		return err
	}

	err := h.uc.AuthUc.CreateUser(ctx, param)
	if err != nil {
		return err
	}
	return response.SuccessOK(c, nil, "Success to create new account")
}

// createPassword godoc
// @Summary      create user password
// @Description  create user password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body  auth.CreateUserPasswordRequest  true  "CreateUserPasswordRequest"
// @Success      200  {object}  response.Response{data=auth.CreateUserPasswordResponse}  "success create user"
// @Failure      400  {object}  response.ErrorResponse      "validation/bind error"
// @Failure      409  {object}  response.ErrorResponse      "email already exists"
// @Failure      500  {object}  response.ErrorResponse      "internal error"
// @Router       /auth/create-password [post]  // <-- adjust to your actual route
func (h *handler) createPassword(c echo.Context) error {
	ctx := c.Request().Context()
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(h.createPassword))
	defer span.End()

	param := auth.CreateUserPasswordRequest{}
	if err := c.Bind(&param); err != nil {
		return err
	}
	if err := param.Validate(); err != nil {
		return err
	}
	resp, err := h.uc.AuthUc.CreatePassword(ctx, param)
	if err != nil {
		return err
	}
	return response.SuccessOK(c, resp, "Success to create new password")
}

// createPin godoc
// @Summary      create user pin
// @Description  create user pin
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body  auth.CreateUserPinRequest  true  "CreateUserPinRequest"
// @Success      200  {object}  response.Response{data=object}  "success create user"
// @Failure      400  {object}  response.ErrorResponse      "validation/bind error"
// @Failure      409  {object}  response.ErrorResponse      "email already exists"
// @Failure      500  {object}  response.ErrorResponse      "internal error"
// @Router       /auth/create-pin [post]  // <-- adjust to your actual route
func (h *handler) createPin(c echo.Context) error {
	ctx := c.Request().Context()
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(h.createPin))
	defer span.End()

	param := auth.CreateUserPinRequest{}
	if err := c.Bind(&param); err != nil {
		return err
	}
	if err := param.Validate(); err != nil {
		return err
	}
	err := h.uc.AuthUc.CreatePin(ctx, param)
	if err != nil {
		return err
	}
	return response.SuccessOK(c, nil, "Success to create new pin")
}

// login godoc
// @Summary      Login
// @Description  Login to app
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body  auth.LoginRequest  true  "LoginRequest"
// @Success      200  {object}  response.Response{data=auth.LoginResponse}  "success login user"
// @Failure      400  {object}  response.ErrorResponse      "validation/bind error"
// @Failure      409  {object}  response.ErrorResponse      "email already exists"
// @Failure      500  {object}  response.ErrorResponse      "internal error"
// @Router       /auth/login [post]  // <-- adjust to your actual route
func (h *handler) login(c echo.Context) (err error) {
	ctx := c.Request().Context()
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(h.login))
	defer span.End()
	//convert and build request
	param := auth.LoginRequest{}
	if err := c.Bind(&param); err != nil {
		return err
	}
	if err := param.Validate(); err != nil {
		return err
	}
	resp, err := h.uc.AuthUc.Login(ctx, param)
	if err != nil {
		return err
	}
	return response.SuccessOK(c, resp, "success login user")
}
