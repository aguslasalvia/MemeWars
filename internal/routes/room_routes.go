package routes

import (
	"memewars/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RoomRoutes(r *gin.RouterGroup) {
	h := handlers.RoomHandler()
	rooms := r.Group("/rooms")
	{
		rooms.POST("", h.CreateRoom)
		rooms.GET("/:id", h.GetRoom)
		rooms.GET("/:id/ranking", h.GetRanking)
	}
}
