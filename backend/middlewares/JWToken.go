package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserInfo struct {
	ID string
	Name string
	Email string
	Role string
	jwt.RegisteredClaims
}

func SecretKey() []byte {
	var secretKey = []byte(os.Getenv("SECRETKEY"))

	return secretKey
}

func CreateToken(id, name, email, role string) (string, error) {
	claims := UserInfo{
		ID: id,
		Name: name,
		Email: email,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SecretKey()); if err != nil {
		return "", err
	}

	return tokenString, err
}

func ParseToken(tokenString string) (*UserInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserInfo{}, func(t *jwt.Token) (interface{}, error) {
		return SecretKey(), nil
	}); if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserInfo); ok && token.Valid{
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}