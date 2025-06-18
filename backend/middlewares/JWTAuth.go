package middlewares

import (
	"net/http"
	"strings"

	"slices"

	"github.com/labstack/echo/v4"
)

func Authentication(requiredRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"Message": "Authorization header required"})
			}

			tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, echo.Map{"Message": "Invalid authorization format"})
			}

			claims, err := ParseToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"Message": "Invalid or expired token"})
			}

			if len(requiredRoles) > 0 {
				roleAllowed := slices.Contains(requiredRoles, claims.Role)
				if !roleAllowed {
					return c.JSON(http.StatusForbidden, echo.Map{"Message": "You don`t have access !"})
				}
			}

			c.Set("user_id", claims.ID)
			c.Set("user_role", claims.Role)

			return next(c)
		}
	}
}