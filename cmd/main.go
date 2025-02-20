package main

import (
	"go-restapi-gin/config"
	"go-restapi-gin/internal/routes"
	"go-restapi-gin/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

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

	// Load Kafka configuration
	kafkaConfig := config.LoadKafkaConfig(cfg)

	// Inisialisasi Kafka Service
	kafkaService, err := services.NewKafkaService(kafkaConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka service: %v", err)
	}
	defer kafkaService.Close() // Close Kafka service when the application finishes

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, db, cfg.JWTSecretKey, kafkaService)

	// Default route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin App is running!",
		})
	})

	// Run server
	router.Run(cfg.ServerAddress)
}
