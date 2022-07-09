package helpers

import (
	"auth-service/conf"
	"auth-service/middlewares"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userData *middlewares.UserData) (result string, err error) {
	claims := middlewares.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "my-auth-service",
			ExpiresAt: time.Now().Add(conf.AppConfig.JWTConfig.JWTExpirationTime).Unix(),
		},
		UID:      userData.UID,
		Username: userData.Username,
		IsAdmin:  userData.IsAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err = token.SignedString([]byte(conf.AppConfig.JWTConfig.JWTSignatureKey))
	return
}
