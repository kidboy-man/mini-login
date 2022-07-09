package helpers

import (
	"auth-service/conf"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	jwt.StandardClaims
	UID      string `json:"uid"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}

type JWTConfig struct {
	JWTPrivateKey     string
	JWTPublicKey      string
	JWTExpirationTime time.Duration
}

type UserData struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}

func GenerateToken(userData *UserData) (result string, err error) {
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "my-auth-service",
			ExpiresAt: time.Now().Add(conf.AppConfig.JWTConfig.JWTExpirationTime).Unix(),
		},
		UID:      userData.UID,
		Username: userData.Username,
		IsAdmin:  userData.IsAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err = token.SignedString([]byte(conf.AppConfig.JWTConfig.JWTPrivateKey))
	return
}
