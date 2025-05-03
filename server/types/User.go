package types

type User struct {
	ID        *int   `json:"id,omitempty"`
	Name      string `json:"username"`
	Email     string `json:"email"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Password  string `json:"password"`
}
