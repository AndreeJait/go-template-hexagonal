package middleware

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/constant"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/AndreeJait/go-template-hexagonal/internal/utils"
	"github.com/AndreeJait/go-utility/errow"
	"github.com/AndreeJait/go-utility/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"strings"
)

func GetUser(c echo.Context) domain.UserRolePermission {
	jwtStr := c.Get(constant.KeyUser)
	user, _ := utils.StringToObject[domain.UserRolePermission](jwtStr.(string))
	return user
}

func MustLogged(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var jwtSecret = cfg.Jwt.Secret
			rawToken := c.Request().Header.Get("Authorization")
			tokens := strings.Split(rawToken, "Bearer ")
			var token = tokens[0]
			if len(strings.TrimSpace(token)) == 0 {
				return errow.ErrInvalidToken
			}
			data, err := jwt.ParseToken[domain.UserRolePermission](token, jwtSecret)
			if err != nil {
				return errors.Wrap(errow.ErrInvalidToken, err.Error())
			}
			c.Set(constant.KeyUser, utils.UnSafeObjectToJsonString(data))
			return next(c)
		}
	}
}
