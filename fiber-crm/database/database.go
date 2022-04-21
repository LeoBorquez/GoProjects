package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBconn *gorm.DB
)
