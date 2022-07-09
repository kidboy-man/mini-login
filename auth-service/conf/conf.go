package conf

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var AppConfig Config

type Config struct {
	DbClient *gorm.DB

	JWTConfig         *JWTConfig
	JWTExpirationTime time.Duration
}

func init() {
	err := godotenv.Load() //Load .env file
	if err != nil {
		panic("fail to load .env file")
	}

	AppConfig.JWTConfig = &JWTConfig{}

	AppConfig.JWTConfig.JWTPrivateKey = os.Getenv("jwt_private_key")
	if AppConfig.JWTConfig.JWTPrivateKey == "" {
		panic("jwt_private_key not set")
	}

	AppConfig.JWTConfig.JWTPublicKey = os.Getenv("jwt_public_key")
	if AppConfig.JWTConfig.JWTPublicKey == "" {
		panic("jwt_public_key not set")
	}

	jwtExpirationTimeStr := os.Getenv("jwt_public_key")
	jwtExpirationTime, _ := strconv.Atoi(jwtExpirationTimeStr)
	if jwtExpirationTime == 0 {
		jwtExpirationTime = 24 * 60 * 60
	}
	AppConfig.JWTExpirationTime = time.Duration(jwtExpirationTime) * time.Second

}
