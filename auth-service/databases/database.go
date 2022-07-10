package database

import (
	"auth-service/conf"
	"auth-service/models"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {

	err := godotenv.Load() //Load .env file
	if err != nil {
		panic(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s port=%s, user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) // Build connection string
	log.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conf.AppConfig.DbClient = conn
	err = conf.AppConfig.DbClient.Debug().AutoMigrate(&models.Auth{}) // Database migration
	if err != nil {
		panic(err)
	}

	if err := conf.AppConfig.DbClient.First(&models.Auth{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		//Insert seed data
		boolFalse := false
		boolTrue := true
		err = conf.AppConfig.DbClient.Model(&models.Auth{}).Create([]*models.Auth{
			{
				UserID:   "esYyQJL2",
				Username: "i_am_admin",
				IsAdmin:  &boolTrue,
				Password: "$2a$14$85BGgoTo9KmjnuH60j/qiOzp.fQ6TzyZZh5Gs9jwNZP7.x9kKT7Me",
			},

			{
				UserID:   "esYyQJL3",
				Username: "i_am_user_1",
				IsAdmin:  &boolFalse,
				Password: "$2a$14$85BGgoTo9KmjnuH60j/qiOzp.fQ6TzyZZh5Gs9jwNZP7.x9kKT7Me",
			},

			{
				UserID:   "esYyQJL5",
				Username: "i_am_user_2",
				IsAdmin:  &boolFalse,
				Password: "$2a$14$85BGgoTo9KmjnuH60j/qiOzp.fQ6TzyZZh5Gs9jwNZP7.x9kKT7Me",
			},
		}).Error
		if err != nil {
			panic(err)
		}
	}

}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return conf.AppConfig.DbClient
}
