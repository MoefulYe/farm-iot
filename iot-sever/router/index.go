package router

import (
	"github.com/MoefulYe/farm-iot/iot-server/handler"
	. "github.com/MoefulYe/farm-iot/iot-server/server"
)

func RegisterRouter() {
	if token := Server.Subscribe(
		"cow/register", 0,
		handler.RegisterHandler,
	); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := Server.Subscribe(
		"cow/login", 0, handler.LoginHandler,
	); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := Server.Subscribe(
		"cow/keep-alive", 0,
		handler.KeepAliveMsgHandler,
	); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := Server.Subscribe("cow/die", 0, handler.DieMsgHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
