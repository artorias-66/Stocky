package main

import (
	"stock-reward-system/src/database"
	"stock-reward-system/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	database.ConnectDB()

	// Set the log format
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// Create a new Gin router
	r := gin.Default()

	// Setup routes
	routes.RegisterRewardRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
