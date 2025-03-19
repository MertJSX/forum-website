package types

type ErrorResponse struct {
	IsError  bool   `json:"error"`
	ErrorMsg string `json:"msg"`
}
