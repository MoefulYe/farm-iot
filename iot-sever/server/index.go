package server

import (
	"github.com/MoefulYe/farm-iot/iot-server/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Server mqtt.Client

func init() {
	opts := config.Conf.NewServerOpts()
	Server = mqtt.NewClient(opts)
	if token := Server.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
