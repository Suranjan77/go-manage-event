package models

type ErrorResponse struct {
	Error     []Error `json:"errors"`
	TimeStamp int64   `json:"timeStamp"`
}

type Error struct {
	Msg string `json:"msg"`
}
