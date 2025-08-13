package auth

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http"
	"github.com/AndreeJait/go-template-hexagonal/internal/usecase"
	"github.com/AndreeJait/go-template-hexagonal/internal/usecase/auth"
	"github.com/AndreeJait/go-utility/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	route *echo.Group
	uc    *usecase.UseCase
}

func NewAuthHandler(route *echo.Group, uc *usecase.UseCase) http.Handler {
	return &handler{
		route: route,
		uc:    uc,
	}
}

func (h *handler) Handle() {
	group := h.route.Group("/auth")

	group.POST("/email", h.getUserByEmail)
}

func (h *handler) getUserByEmail(c echo.Context) (err error) {
	ctx := c.Request().Context()
	param := auth.GetUserByEmailRequest{}
	if err := c.Bind(&param); err != nil {
		return err
	}
	user, err := h.uc.AuthUc.GetUserByEmail(ctx, param)
	if err != nil {
		return err
	}
	return response.SuccessOK(c, user)
}
