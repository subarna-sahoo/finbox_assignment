package utils

import (
	"finbox_assignment/database"
	"finbox_assignment/models"
	"log"
)

func AutoMigrate() {
	db := database.GetDB()

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
