package models

type Resp[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func NewResp[T any](code int, msg string, data T) Resp[T] {
	return Resp[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

type Msg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func MsgOnly(code int, msg string) Msg {
	return Msg{
		Code: code,
		Msg:  msg,
	}
}
