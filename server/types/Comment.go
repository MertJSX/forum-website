package types

type Comment struct {
	ID        *int   `json:"id,omitempty"`
	UserId    string `json:"user_id"`
	ForumId   string `json:"forum_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}
