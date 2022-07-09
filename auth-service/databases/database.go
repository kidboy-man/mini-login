package database

import (
	"auth-service/conf"
	"auth-service/models"
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //database

func InitDB() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s port=%s, user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) // Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conf.AppConfig.DbClient = conn
	conf.AppConfig.DbClient.Debug().AutoMigrate(&models.Auth{}) // Database migration
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
