package main

import (
	"memewars/internal/config"
	"memewars/internal/db"
	"memewars/internal/routes"
)

func main() {
	// Run the connection && migration function
	db.Connect()

	// Loads all the existing routes of the domain
	r := routes.SetupRoutes()

	// Start the server using .env PORT or by default the 4040
	r.Run(":" + config.GetEnv("PORT", "4040"))
}
