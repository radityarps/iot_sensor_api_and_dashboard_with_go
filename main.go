package main

import (
	"iot_api_with_go/database"
	"iot_api_with_go/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	 r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	// Register routes
	routes.SensorStatusRoutes(r)

	r.Run(":8080")
}
