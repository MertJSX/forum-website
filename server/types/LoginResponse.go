package types

type LoginResponse struct {
	Msg   string `json:"msg"`
	Token string `json:"token"`
}
