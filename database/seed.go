package database

import (
	"log"

	"gorm.io/gorm"
)

// Seeder is a function type that seeds data into the database
type Seeder func(*gorm.DB) error

// Seed runs all seed functions
func Seed(db *gorm.DB, seeders ...Seeder) error {
	log.Println("🌱 Seeding database...")

	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return err
		}
	}

	log.Println("✅ Database seeding completed")
	return nil
}
