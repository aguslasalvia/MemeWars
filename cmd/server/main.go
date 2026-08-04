package main

import (
	"memewars/internal/config"
	"memewars/internal/db"
	"memewars/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Run the connection && migration function
	db.Connect()

	// Loads all the existing routes of the domain
	r := routes.SetupRoutes()

	gin.SetMode(gin.ReleaseMode)

	// Start the server using .env PORT or by default the 4040
	r.Run(":" + config.GetEnv("PORT", "4040"))
}
