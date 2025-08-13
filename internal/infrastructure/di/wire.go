package di

import (
	"context"
	"github.com/AndreeJait/go-template-hexagonal/internal/constant"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/db"
	"github.com/AndreeJait/go-utility/configw"
	"github.com/AndreeJait/go-utility/gracefull"
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/labstack/echo/v4"
	"os"
)

func Wire(ctx context.Context) (*echo.Echo, *gracefull.GracefulShutDown, loggerw.Logger, error) {

	var appMode = os.Getenv("APP_ENV")
	var logFormat = loggerw.JSONFormatter

	if constant.AppMode(appMode) == constant.DevelopmentMode {
		logFormat = loggerw.TextFormatter
	}

	l, err := loggerw.New(&loggerw.Option{
		Formatter: logFormat,
	})
	if err != nil {
		panic(err)
	}
	// handle graceful
	g := gracefull.NewGracefulShutdown(l)

	cfgInit := configw.New[config.Config](config.MapFiles,
		config.ConfigMode[appMode])

	cfg, err := cfgInit.LoadConfig()
	if err != nil {
		l.Errorf(ctx, err, "load config error: %v", err)
		return nil, g, l, err
	}

	pg, err := db.NewPostgres(ctx, cfg.DB)
	if err != nil {
		l.Fatalf(ctx, "new postgres error: %v", err)
		return nil, g, l, err
	}
	g.AddFunc("postgres", func() error {
		pg.Close()
		l.Info(ctx, "postgresql disconnected")
		return nil
	})

	srv := NewWire(cfg, pg, l)
	e := srv.initHandler()
	g.AddFunc("http", func() error {
		l.Info(ctx, "wire shutting down")
		return e.Shutdown(ctx)
	})
	return e, g, l, nil
}
