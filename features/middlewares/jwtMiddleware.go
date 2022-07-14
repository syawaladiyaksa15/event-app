package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("SecretJWT")),
	})
}

func CreateToken(userId int, avatarUrl string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["avatarUrl"] = avatarUrl
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SecretJWT")))
}

func ExtractToken(e echo.Context) (int, string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		avatarUrl := claims["avatarUrl"].(string)
		return int(userId), avatarUrl, nil
	}
	return 0, "Avatar link not found", fmt.Errorf("token invalid")
}
