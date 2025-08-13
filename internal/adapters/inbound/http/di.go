package http

import (
	"fmt"
	"github.com/AndreeJait/go-template-hexagonal/internal/constant"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/AndreeJait/go-utility/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func NewEcho(cfg *config.Config, log loggerw.Logger) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	// setup port
	e.Server.Addr = fmt.Sprintf("%s:%s", cfg.Service.Host, cfg.Service.Port)

	isDevelopment := config.IsDevelopment()
	e.HTTPErrorHandler = response.CustomHttpErrorHandler(log, constant.ErrorMap, isDevelopment)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	//e.Use(middleware.Recover())
	e.Use(loggerw.LoggerWithRequestID(log, isDevelopment))

	// health
	e.GET("/healthz", healthz)

	return e
}

func healthz(c echo.Context) error {
	return response.SuccessOK(c, echo.Map{"status": "ok"})
}
