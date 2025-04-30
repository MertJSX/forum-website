package types

type Forum struct {
	ID          *int   `json:"id,omitempty"`
	UserId      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Content     string `json:"content"`
}
