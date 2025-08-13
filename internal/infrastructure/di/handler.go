package di

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/inbound/http/auth"
	"github.com/labstack/echo/v4"
)

func (s wire) initHandler() *echo.Echo {
	// register repo
	s.initRepository()
	// register use case here
	s.initUseCase()

	e := http.NewEcho(s.cfg, s.log)

	groupV1 := e.Group("/api/v1")
	var handlers = []http.Handler{
		auth.NewAuthHandler(groupV1, s.uc),
	}

	for _, h := range handlers {
		h.Handle()
	}
	return e
}
