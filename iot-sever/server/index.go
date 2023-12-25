package server

import (
	"github.com/MoefulYe/farm-iot/iot-server/config"
	"github.com/MoefulYe/farm-iot/iot-server/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

var Server mqtt.Client

func init() {
	opts := config.Conf.NewServerOpts()
	Server = mqtt.NewClient(opts)
	if token := Server.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf(token.Error().Error())
	}
	logger.Logger.Infow("init mqtt client")
}
