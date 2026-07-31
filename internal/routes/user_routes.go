package routes

import (
	"memewars/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	h := handlers.UserHandler()
	users := r.Group("/users")
	{
		users.POST("", h.Create)
		users.GET("/:name", h.GetUserByName)
	}
}
