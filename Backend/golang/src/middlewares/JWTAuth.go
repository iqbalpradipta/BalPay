package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
    return middleware.JWTWithConfig(middleware.JWTConfig{
        SigningMethod: "HS256",
        SigningKey:    []byte(os.Getenv("SECRET_KEY_JWT")),
    })
}

func ExtractToken(c echo.Context) (int, string) {
    headerData := c.Request().Header.Get("Authorization")
    dataAuth := strings.Split(headerData, " ")
    token := dataAuth[len(dataAuth)-1]
    datajwt, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("SECRET_KEY_JWT")), nil
    })

    if datajwt.Valid {
        claims := datajwt.Claims.(jwt.MapClaims)
        userId, userIdOk := claims["id"].(float64)
        role, roleOk := claims["role"].(string)
        if userIdOk && roleOk {
            return int(userId), role
        }
        if !userIdOk {
            fmt.Println("UserID is missing or not a float64")
        }
        if !roleOk {
            fmt.Println("Role is missing or not a string")
        }
    }
    return -1, ""
}



func Authentication(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, role := ExtractToken(c)
			for _, r := range roles {
				if r == role {
					return next(c)
				}
			}
			return echo.ErrForbidden
		}
	}
}

