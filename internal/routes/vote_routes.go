package routes

import (
	"memewars/internal/handlers"

	"github.com/gin-gonic/gin"
)

func VoteRoutes(r *gin.RouterGroup) {
	h := handlers.VoteHanlder()
}
