package router

import (
	"github.com/MoefulYe/farm-iot/iot-server/handler"
	. "github.com/MoefulYe/farm-iot/iot-server/logger"
	. "github.com/MoefulYe/farm-iot/iot-server/server"
)

func RegisterRouter() {
	if token := Server.Subscribe(
		"cow/register",
		0,
		handler.RegisterHandler,
	); token.Wait() && token.Error() != nil {
		Logger.Fatal(token.Error().Error())
	}
	if token := Server.Subscribe(
		"cow/login",
		0,
		handler.LoginHandler,
	); token.Wait() && token.Error() != nil {
		Logger.Fatal(token.Error().Error())
	}
	if token := Server.Subscribe(
		"cow/heartbeat",
		0,
		handler.HeartBeatHandler,
	); token.Wait() && token.Error() != nil {
		Logger.Fatal(token.Error().Error())
	}
	if token := Server.Subscribe(
		"cow/die",
		0,
		handler.DieHandler,
	); token.Wait() && token.Error() != nil {
		Logger.Fatal(token.Error().Error())
	}
}
