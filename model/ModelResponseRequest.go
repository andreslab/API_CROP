package model

type ResponseRequestModel struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ResponseRequestIdModel struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	ID   int64  `json:"id"`
}
