package handlers

import (
	"memewars/internal/models"
	"memewars/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type voteHandler struct {
	service *services.VoteService
}

func VoteHandler() *voteHandler {
	return &voteHandler{
		service: &services.VoteService{},
	}
}

func (vh *voteHandler) Create(ctx *gin.Context) {
	var req models.Vote

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v, err := vh.service.GiveVoteToMeme(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, v)
}
