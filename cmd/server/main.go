package main

import (
	"memewars/internal/config"
	"memewars/internal/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run(config.GetEnv("PORT", "4040"))
}
