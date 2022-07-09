package helpers

import (
	"auth-service/conf"
	"encoding/base64"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	IssuerJWT = "github.com"
)

type JWTClaims struct {
	UserData
	jwt.StandardClaims
}

type JWTConfig struct {
	JWTPrivateKey     string
	JWTPublicKey      string
	JWTExpirationTime time.Duration
}

type UserData struct {
	UID     string `json:"uid"`
	IsAdmin bool   `json:"isAdmin"`
}

type UserDataAdmin struct {
	UID     string `json:"uid"`
	IsAdmin bool   `json:"isAdmin"`
	RoleIDs []int  `json:"roleIDs"`
}

type JWTClaimsAdmin struct {
	UserDataAdmin
	jwt.StandardClaims
}

func GenerateToken(userData *UserData) (result string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"uid":     userData.UID,
		"isAdmin": userData.IsAdmin,
		"iss":     IssuerJWT,
		"sub":     IssuerJWT,
		"exp":     time.Now().Add(conf.AppConfig.JWTConfig.JWTExpirationTime).Unix(),
		"iat":     time.Now().Unix(),
	})

	privateKey, err := base64.StdEncoding.DecodeString(conf.AppConfig.JWTConfig.JWTPrivateKey)
	if err != nil {
		return
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return
	}

	result, err = token.SignedString(key)
	return
}
