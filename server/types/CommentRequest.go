package types

type CommentRequest struct {
	Token   string  `json:"token"`
	Comment Comment `json:"comment"`
}
