package types

type Post struct {
	ID        *int   `json:"id,omitempty"`
	UserId    int    `json:"userid"`
	Upvotes   int    `json:"upvotes"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	Content   string `json:"content"`
}
