package handlers

import (
	"memewars/internal/models"
	"memewars/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type roomHandler struct {
	service *services.RoomService
}

func RoomHandler() *roomHandler {
	return &roomHandler{}
}

func (rh *roomHandler) CreateRoom(ctx *gin.Context) {
	var req models.CreateRoomRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := rh.service.CreateRoom(req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating a room"})
	}

	ctx.JSON(http.StatusCreated, room)
}

func (rh *roomHandler) GetRoom() {

}

func (rh *roomHandler) UploadMeme() {

}

func (rh *roomHandler) GetRanking() {

}
