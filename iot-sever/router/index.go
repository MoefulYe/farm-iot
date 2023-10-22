package router

import (
	"github.com/MoefulYe/farm-iot/iot-server/handler"
	. "github.com/MoefulYe/farm-iot/iot-server/server"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func RegisterRouter(server mqtt.Client) {
	Server.Subscribe("cow/register", 0, handler.RegisterHandler)
	Server.Subscribe("cow/login", 0, handler.LoginHandler)
	Server.Subscribe("cow/keep-alive", 0, handler.KeepAliveMsgHandler)
}
