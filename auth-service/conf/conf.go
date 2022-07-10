package conf

import (
	"log"
	"os"
	"strconv"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var AppConfig Config

type Config struct {
	DbClient       *gorm.DB
	UserServiceURL string

	JWTConfig *JWTConfig
}

type JWTConfig struct {
	JWTSignatureKey   string
	JWTExpirationTime time.Duration
}

func init() {
	err := godotenv.Load() //Load .env file
	if err != nil {
		panic("fail to load .env file")
	}

	beego.BConfig.RunMode = os.Getenv("beego_runmode")
	log.Println("beego.BConfig.RunMode", beego.BConfig.RunMode)

	AppConfig.JWTConfig = &JWTConfig{}

	AppConfig.JWTConfig.JWTSignatureKey = os.Getenv("jwt_signature_key")
	if AppConfig.JWTConfig.JWTSignatureKey == "" {
		panic("jwt_signature_key not set")
	}

	AppConfig.UserServiceURL, _ = beego.AppConfig.String("user_service_url")
	log.Println("AppConfig.UserServiceURL", AppConfig.UserServiceURL)

	jwtExpirationTimeStr := os.Getenv("jwt_expiration_time")
	jwtExpirationTime, _ := strconv.Atoi(jwtExpirationTimeStr)
	if jwtExpirationTime == 0 {
		jwtExpirationTime = 24 * 60 * 60
	}
	AppConfig.JWTConfig.JWTExpirationTime = time.Duration(jwtExpirationTime) * time.Second

}
