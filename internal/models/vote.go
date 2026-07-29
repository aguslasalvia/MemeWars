package models

type Vote struct {
	ID     string `json:"id"`
	MemeID string `json:"meme_id"`
	UserID string `json:"user_id"`
	Value  int    `json:"vote"`
}
