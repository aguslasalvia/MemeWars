package routes

import (
	"memewars/internal/handlers"

	"github.com/gin-gonic/gin"
)

func VoteRoutes(r *gin.RouterGroup) {
	h := handlers.VoteHandler()
	votes := r.Group("/vote")
	{
		votes.POST("", h.Create)
	}
}
