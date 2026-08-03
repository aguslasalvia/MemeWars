package routes

import (
	"strings"

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

	// Uploaded meme images, saved by the meme service under ./static/memes
	r.Static("/static/memes", "./static/memes")

	// Frontend build output (bun run build -> frontend/dist -> ../static/dist)
	r.Static("/assets", "./static/dist/assets")
	r.StaticFile("/favicon.svg", "./static/dist/favicon.svg")
	r.StaticFile("/", "./static/dist/index.html")

	group := r.Group("/api/v1")
	{
		RoomRoutes(group)
		UserRoutes(group)
		MemeRoutes(group)
		VoteRoutes(group)
	}

	// SPA fallback so client-side routes (e.g. /room/1) survive a refresh
	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Status(404)
			return
		}
		c.File("./static/dist/index.html")
	})

	return r
}
