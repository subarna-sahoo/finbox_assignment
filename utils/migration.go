package utils

import (
	"finbox_assignment/models"
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.FeatureFlag{},
		&models.Dependency{},
		&models.Client{},
		&models.ClientFeatureFlag{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
