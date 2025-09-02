package di

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http/auth"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http/swagger"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http/user"
	"github.com/labstack/echo/v4"
)

func (s wire) initHandler() *echo.Echo {
	// register repo
	s.initRepository()
	// register use case here
	s.initUseCase()

	e := http.NewEcho(s.cfg, s.log)

	swagger.MountSwagger(e, swagger.SwaggerCfg{
		Host: s.cfg.Service.Host,
		Port: s.cfg.Service.Port,
	})

	groupV1 := e.Group("/api/v1")
	var handlers = []http.Handler{
		auth.NewAuthHandler(s.cfg, groupV1, s.uc),
		user.NewUserHandler(s.cfg, groupV1, s.uc),
	}

	for _, h := range handlers {
		h.Handle()
	}
	return e
}
