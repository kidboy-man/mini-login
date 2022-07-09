package conf

import "gorm.io/gorm"

var AppConfig Config

type Config struct {
	DbClient *gorm.DB
}
