package main

import (
	"log"
	"go-restapi-gin/config"
	"go-restapi-gin/internal/routes"

    "github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ensure database connection is closed properly
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection:", err)
	}
	defer sqlDB.Close()

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, db)

	// Default route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin App is running!",
		})
	})

	// Run server
	router.Run(cfg.ServerAddress)
}

