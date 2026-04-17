package main

import (
	"go-pagination/config"
	"go-pagination/models"
	"go-pagination/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Connect to Database
	config.ConnectDatabase()

	// 2. Run Migrations
	config.DB.AutoMigrate(&models.Product{})

	// 3. Seed Data
	models.SeedProducts(config.DB)

	// 4. Setup Router
	r := gin.Default()

	// 5. Register Routes
	routes.SetupRoutes(r)

	// 6. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
