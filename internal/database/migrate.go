package database

import (
	"github.com/aminramezanipoor/url-shortener/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.URL{},
	)
}