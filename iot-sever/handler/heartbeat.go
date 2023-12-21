package handler

import (
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow/heartbeat"
	"github.com/MoefulYe/farm-iot/iot-server/stream"
	"github.com/MoefulYe/farm-iot/iot-server/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"time"
)

func HeartBeatHandler(server mqtt.Client, msg mqtt.Message) {
	keepAlive := new(heartbeat.HeartBeat)
	if err := proto.Unmarshal(msg.Payload(), keepAlive); err != nil {
		return
	}
	claims, err := utils.JWTParse(keepAlive.GetToken())
	if err != nil {
		return
	}
	id := uuid.MustParse(claims.Id)
	timestamp, err := time.Parse(time.RFC3339, keepAlive.GetTimestamp())
	if err != nil {
		return
	}
	data := &stream.HeartBeat{
		Uuid:      id,
		Timestamp: timestamp,
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
	stream.Input() <- data
}
