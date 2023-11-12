package models

import (
	"github.com/google/uuid"
	"time"
)

type ResponseList struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Resp[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}
type RegisterList struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}

type Uuid struct {
	Uuid uuid.UUID `json:"uuid"`
}
type Info struct {
	Id     string     `json:"id"`
	BornAt time.Time  `json:"born_at"`
	DeadAt *time.Time `json:"dead_at,omitempty"`
	Reason *string    `json:"reason,omitempty"`
}

type KeepalivePackage struct {
	Id    string `json:"id"`
	Value string `json:"value"`
	Time  string `json:"time"`
}
