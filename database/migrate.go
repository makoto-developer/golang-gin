package database

import (
	"log"

	"gorm.io/gorm"
)

// Migrate runs auto migration for all models
func Migrate(db *gorm.DB, models ...interface{}) error {
	log.Println("🔄 Running database migrations...")

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	log.Println("✅ Database migrations completed")
	return nil
}
