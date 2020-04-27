package model

// Request is ...
// リクエストパラメーター
type Request struct {
	UserID int    `json:"userID"`
	Name   string `json:name"`
}

// Response is ...
// レスポンスパラメーター
type Response struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}
