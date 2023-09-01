package db

import "gorm.io/gorm"

type mysqlDb struct {
	db *gorm.DB
}
