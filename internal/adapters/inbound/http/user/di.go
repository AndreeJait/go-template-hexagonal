package user

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http/common/middleware"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/AndreeJait/go-template-hexagonal/internal/usecase"
	"github.com/AndreeJait/go-utility/response"
	"github.com/AndreeJait/go-utility/tracer"
	"github.com/labstack/echo/v4"
)

type handler struct {
	route *echo.Group
	uc    *usecase.UseCase
	cfg   *config.Config
}

func NewUserHandler(cfg *config.Config, route *echo.Group, uc *usecase.UseCase) http.Handler {
	return &handler{
		route: route,
		uc:    uc,
		cfg:   cfg,
	}
}

func (h *handler) Handle() {
	groupPrivate := h.route.Group("/user")
	groupPrivate.Use(middleware.MustLogged(h.cfg))

	//groupPrivate.GET("/me", h.me, middleware.CheckUserCan(domain.UserCanCreateCustomer))
	groupPrivate.GET("/me", h.me)
}

// me godoc
// @Summary      get user info
// @Description  get user info
// @Tags         User
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  response.Response{data=domain.UserSimplified}  "success create user"
// @Failure      401  {object}  response.ErrorResponse      "forbidden"
// @Failure      500  {object}  response.ErrorResponse      "internal error"
// @Router       /user/me [get]  // <-- adjust to your actual route
func (h *handler) me(c echo.Context) error {
	ctx := c.Request().Context()
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(h.me))
	defer span.End()

	userLogged := middleware.GetUser(c)
	var user domain.UserSimplified
	var err error
	user, err = h.uc.UserUc.GetUserById(ctx, userLogged.ID)
	if err != nil {
		return err
	}
	return response.SuccessOK(c, user, "success get user")
}
