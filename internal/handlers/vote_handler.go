package handlers

type voteHandler struct{}

func VoteHandler() *voteHandler {
	return &voteHandler{}
}
