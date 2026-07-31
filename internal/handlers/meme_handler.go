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

	roomP := ctx.Param("room_id")
	memeP := ctx.Param("room_id")

	if roomP == "" || memeP == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "RoomID and MemeID must be provided"})
	}

	roomID, room_err := strconv.ParseUint(roomP, 10, 32)
	memeID, meme_err := strconv.ParseUint(memeP, 10, 32)

	if room_err != nil || meme_err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	m, err := mh.service.FindMemeFromRoom(uint(memeID), uint(roomID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
	}

	ctx.JSON(http.StatusOK, m)

}
