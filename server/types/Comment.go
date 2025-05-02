package types

type Comment struct {
	ID        *int   `json:"id,omitempty"`
	UserId    *int   `json:"user_id"`
	PostId    string `json:"post_id"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
}
