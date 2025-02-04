package database

import (
	"book/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Init() (*gorm.DB, error) {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Book{})
	}
	return db, err
}
