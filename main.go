package main

import (
	"log"

	"finbox_assignment/database"
	"finbox_assignment/utils"
)

func main() {
	// Initialize database connection
	_ = database.GetDB()

	// Run migrations
	utils.AutoMigrate()

	log.Println("Migrations completed successfully")
}
