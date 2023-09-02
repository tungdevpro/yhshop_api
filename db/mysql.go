package db

import (
	"coffee_api/configs"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func CreateMysqlDB(cfg *configs.Configuration) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DBConnectionURL), &gorm.Config{})
	return db, err
}
