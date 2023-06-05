package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/iqbalpradipta/BalPay/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.SECRET_JWT),
	})
}

func CreateToken(id int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractToken(c echo.Context) (int, string) {
	user := c.Get("user").(*jwt.Token)

	if user.Valid{
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		role := claims["role"].(string)
		return int(id), role
	}

	return 0, ""
}