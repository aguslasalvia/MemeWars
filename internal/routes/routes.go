package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	r.Static("/static", "./static")

	group := r.Group("/api/v1")
	{
		RoomRoutes(group)
		UserRoutes(group)
		MemeRoutes(group)
		VoteRoutes(group)
	}

	return r
}
