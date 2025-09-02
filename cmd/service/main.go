package main

import (
	"context"
	"errors"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/di"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Package main
//
//	@title           Hexagonal Template API
//	@version         1.0
//	@description     Hexagonal template service (Hexagonal, Echo v4, Postgres).
//
//	@contact.name    Dev Team
//	@contact.email   dev@example.com
//
//	@BasePath        /api/v1
//
//	@securityDefinitions.apikey BearerAuth
//	@in header
//	@name Authorization
//	@description  Add prefix: "Bearer <token>"
//
//	@securityDefinitions.basic BasicAuth
func main() {
	if os.Getenv("APP_ENV") == "" {
		_ = os.Setenv("APP_ENV", "development")
	}
	var ctx = context.Background()
	e, g, log, err := di.Wire(ctx)
	if err != nil {
		panic(err)
	}

	go func() {
		addr := e.Server.Addr
		if addr == "" {
			addr = ":8080"
		}
		log.Info(ctx, "listening on ", addr)
		if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(ctx, err, "server stopped:")
		}
	}()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	g.ShutdownAll(shutdownCtx)
}
