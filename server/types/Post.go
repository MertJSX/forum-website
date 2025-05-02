package types

type Post struct {
	ID        *int   `json:"id,omitempty"`
	UserId    int    `json:"userid"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	Content   string `json:"content"`
}
