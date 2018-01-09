package model

type ResponseRequestModel struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ResponseRequestUserModel struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	IDUser int64  `json:"iduser"`
}
