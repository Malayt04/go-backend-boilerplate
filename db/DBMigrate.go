package database

import (
	"github.com/go-backend/models"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) error {

	return db.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Ticket{},
	)

}