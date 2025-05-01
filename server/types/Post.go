package types

type Post struct {
	ID          *int   `json:"id,omitempty"`
	UserId      string `json:"userid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Content     string `json:"content"`
}
