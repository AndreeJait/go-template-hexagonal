package middleware

import (
	"encoding/base64"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/AndreeJait/go-utility/errow"
	"github.com/labstack/echo/v4"
	"strings"
)

func BasicAuthLogged(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")
			if header == "" || !strings.HasPrefix(header, "Basic ") {
				return errow.ErrForbidden
			}

			// Decode base64(username:password)
			payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(header, "Basic "))
			if err != nil {
				return errow.ErrForbidden
			}

			// Split into username:password
			parts := strings.SplitN(string(payload), ":", 2)
			if len(parts) != 2 {
				return errow.ErrForbidden
			}
			username, password := parts[0], parts[1]

			// Compare with config
			if username != cfg.InternalApi.Username || password != cfg.InternalApi.Password {
				return errow.ErrForbidden
			}

			// Add user info to context if needed
			c.Set("auth.user", username)

			return next(c)
		}
	}
}
