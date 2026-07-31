package handlers

import (
	"memewars/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type memeHandler struct {
	service *services.MemeService
}

func MemeHandler() *memeHandler {
	return &memeHandler{}
}

func (mh *memeHandler) GetMemeFromRoom(ctx *gin.Context) {

	roomParam := ctx.Param("room_id")
	memeParam := ctx.Param("room_id")

	if roomParam == "" || memeParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "RoomID and MemeID must be provided"})
	}

	roomID, roomErr := strconv.ParseUint(roomParam, 10, 32)
	memeID, memeErr := strconv.ParseUint(memeParam, 10, 32)

	if roomErr != nil || memeErr != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	m, err := mh.service.FindMemeFromRoom(uint(memeID), uint(roomID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}

	ctx.JSON(http.StatusOK, m)

}

func (mh *memeHandler) GetMemesFromRoom(ctx *gin.Context) {
	roomParam := ctx.Param("room_id")

	if roomParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "RoomID must be provided"})
	}

	roomID, roomErr := strconv.ParseUint(roomParam, 10, 32)
	if roomErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": roomErr})
	}

	ms := mh.service.FindAllMemesFromRoom(uint(roomID))
	ctx.JSON(http.StatusOK, ms)
}
