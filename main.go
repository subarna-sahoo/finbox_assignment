package main

import (
	"finbox_assignment/database"
	"finbox_assignment/routers"
	"finbox_assignment/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db := database.GetDB()

	// Run migrations
	utils.AutoMigrate(db)

	// Setup Gin router
	r := gin.Default()

	// Setup routes
	routers.SetupRoutes(r)

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
