package middleware

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/constant"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
	"github.com/labstack/echo/v4"
)

func CheckUserCan(permissions ...domain.UserPermissionsName) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := GetUser(c)
			var mapPermissions = map[domain.UserPermissionsName]bool{}

			// filled user permission
			for _, permission := range user.Permission {
				mapPermissions[permission.Name] = true
			}

			// check if user have the permission
			for _, permission := range permissions {
				if mapPermissions[permission] {
					return next(c)
				}
			}
			return constant.ErrUserDidNotHaveAccessToThisFeature
		}
	}
}
