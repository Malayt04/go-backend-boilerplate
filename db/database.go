package database

import (
	"fmt"

	"github.com/go-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(env *config.Env, DBMigrator func(db *gorm.DB) error) *gorm.DB {

	uri := fmt.Sprint(
		`host=%s dbname=%s password=%s sslmode=%s port=5432`,
		env.DB_HOST,
		env.DB_NAME,
		env.DB_PASSWORD,
		env.DB_SSLMODE,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

	return db
}