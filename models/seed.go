package models

import (
	"gorm.io/gorm"
)

// SeedAlbums seeds initial album data
func SeedAlbums(db *gorm.DB) error {
	albums := []Album{
		{Title: "Hammerhead", Artist: "THE OFFSPRING", Price: 25.05, Tax: 0.1},
		{Title: "Shake It Off", Artist: "Taylor Swift", Price: 23.14, Tax: 0.1},
		{Title: "mysterious love", Artist: "Miho Komatsu", Price: 18.88, Tax: 0.1},
	}

	// Check if data already exists
	var count int64
	if err := db.Model(&Album{}).Count(&count).Error; err != nil {
		return err
	}

	// Only seed if table is empty
	if count == 0 {
		if err := db.Create(&albums).Error; err != nil {
			return err
		}
	}

	return nil
}
