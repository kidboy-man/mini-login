package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"user-service/conf"
	"user-service/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {

	err := godotenv.Load() //Load .env file
	if err != nil {
		log.Println("error get env ", err)
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
	err = conf.AppConfig.DbClient.Debug().AutoMigrate(&models.User{}) // Database migration
	if err != nil {
		panic(err)
	}

	if err := conf.AppConfig.DbClient.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		//Insert seed data
		err = conf.AppConfig.DbClient.Model(&models.User{}).Create([]*models.User{
			{
				ID:       "esYyQJL2",
				FullName: "Admin",
				Email:    "admin@example.com",
			},

			{
				ID:       "esYyQJL3",
				FullName: "Regular User 1",
				Email:    "user1@example.com",
			},

			{
				ID:       "esYyQJL5",
				FullName: "Regular User 2",
				Email:    "user2@example.com",
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
