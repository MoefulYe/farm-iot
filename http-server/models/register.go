package models

type RegisterReq struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}
