package services

import (
	"errors"
	"memewars/internal/db"
	"memewars/internal/models"
)

type VoteService struct{}

func (vs *VoteService) GiveVoteToMeme(vote *models.Vote) (*models.Vote, error) {
	var m models.Meme
	if err := db.DB.First(&m, vote.MemeID).Error; err != nil {
		return nil, errors.New("Meme not found")
	}

	var eVote models.Vote
	err := db.DB.
		Where("meme_id = ? AND user_id = ?", vote.MemeID, vote.UserID).
		First(&eVote).Error
	if err == nil {
		return nil, errors.New("Vote already set")
	}

	if err := db.DB.Create(vote).Error; err != nil {
		return nil, err
	}

	return vote, nil
}
