package handlers

import (
	"memewars/internal/models"
	"memewars/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type memeHandler struct {
	service *services.MemeService
}

func MemeHandler() *memeHandler {
	return &memeHandler{
		service: &services.MemeService{},
	}
}

func (mh *memeHandler) UploadMeme(ctx *gin.Context) {
	roomID, err := strconv.ParseUint(ctx.Param("room_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid room id"})
		return
	}

	var req models.CreateMemeRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meme, err := mh.service.CreateMeme(uint(roomID), req.UserID, req.Text, req.Image)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, meme)

}

func (mh *memeHandler) GetMemeFromRoom(ctx *gin.Context) {

	roomParam := ctx.Param("room_id")
	memeParam := ctx.Param("meme_id")

	if roomParam == "" || memeParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "RoomID and MemeID must be provided"})
		return
	}

	roomID, roomErr := strconv.ParseUint(roomParam, 10, 32)
	memeID, memeErr := strconv.ParseUint(memeParam, 10, 32)

	if roomErr != nil || memeErr != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	m, err := mh.service.FindMemeFromRoom(uint(memeID), uint(roomID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, m)

}

func (mh *memeHandler) GetMemesFromRoom(ctx *gin.Context) {
	roomParam := ctx.Param("room_id")

	if roomParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "RoomID must be provided"})
		return
	}

	roomID, roomErr := strconv.ParseUint(roomParam, 10, 32)
	if roomErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": roomErr})
		return
	}

	ms := mh.service.FindAllMemesFromRoom(uint(roomID))
	ctx.JSON(http.StatusOK, ms)
}
