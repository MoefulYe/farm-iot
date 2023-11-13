package models

type LoginReq struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}

type Token struct {
	Token string `json:"token"`
}
