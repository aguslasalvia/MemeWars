package routes

import (
	"memewars/internal/handlers"

	"github.com/gin-gonic/gin"
)

func MemeRoutes(r *gin.RouterGroup) {
	h := handlers.MemeHandler()
	memes := r.Group("/memes")
	{
		memes.GET("/:room/:meme_id", h.GetMemeFromRoom)
		memes.GET("/:room_id",h.GetMemesFromRoom)
	}
}
