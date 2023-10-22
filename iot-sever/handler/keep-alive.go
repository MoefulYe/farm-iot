package handler

import (
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow"
	"github.com/MoefulYe/farm-iot/iot-server/stream"
	"github.com/MoefulYe/farm-iot/iot-server/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

func KeepAliveMsgHandler(server mqtt.Client, msg mqtt.Message) {
	keepAlive := new(cow.KeepAliveMsg)
	if err := proto.Unmarshal(msg.Payload(), keepAlive); err != nil {
		return
	}
	claims, err := utils.JWTParse(keepAlive.GetToken())
	if err != nil {
		return
	}
	id := uuid.MustParse(claims.Id)
	data := &stream.KeepAlive{
		Uuid:      id,
		Timestamp: keepAlive.GetTimestamp().AsTime(),
		Weight:    keepAlive.GetWeight(),
		Health:    keepAlive.GetHealth(),
		Geo: struct {
			Latitude  float64
			Longitude float64
		}{
			Latitude:  keepAlive.GetGeo().GetLatitude(),
			Longitude: keepAlive.GetGeo().GetLongitude(),
		},
	}
	stream.StreamInput() <- data
}
