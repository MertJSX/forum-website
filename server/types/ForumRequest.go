package types

type ForumRequest struct {
	Token string `json:"token"`
	Forum Forum  `json:"forum"`
}
