package handlers

import (
	"memewars/internal/models"
	"memewars/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service *services.UserService
}

func UserHandler() *userHandler {
	return &userHandler{}
}

func (uh *userHandler) Create(ctx *gin.Context) {
	var req models.User
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	us, err := uh.service.Create(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, us)
}

func (uh *userHandler) GetUserByName(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name must be inserted"})
	}

	u, err := uh.service.GetUserByName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, u)

}
