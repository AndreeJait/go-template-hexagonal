package swagger

import (
	"fmt"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/swagger/docs"
	_ "github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/swagger/docs"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

type SwaggerCfg struct {
	Host string
	Port string
	// optional:
	Schemes []string // e.g. []string{"http"} or []string{"https"}
}

func MountSwagger(e *echo.Echo, sc SwaggerCfg) {
	// Default scheme
	if len(sc.Schemes) == 0 {
		// Use https automatically on production if you terminate TLS at LB
		if config.IsProduction() {
			sc.Schemes = []string{"https"}
		} else {
			sc.Schemes = []string{"http"}
		}
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", sc.Host, sc.Port)
	docs.SwaggerInfo.Schemes = sc.Schemes
	docs.SwaggerInfo.BasePath = "/api/v1/" // set if you serve behind a prefix

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
