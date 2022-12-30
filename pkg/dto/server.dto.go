package dto

type Response struct {
	Data    any    `json:"data"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  int    `json:"status"`
}
