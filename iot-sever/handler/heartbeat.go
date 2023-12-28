package handler

import (
	"fmt"
	"github.com/MoefulYe/farm-iot/iot-server/constant"
	"github.com/MoefulYe/farm-iot/iot-server/db"
	. "github.com/MoefulYe/farm-iot/iot-server/logger"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow/heartbeat"
	"github.com/MoefulYe/farm-iot/iot-server/stream"
	"github.com/MoefulYe/farm-iot/iot-server/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"google.golang.org/protobuf/proto"
	"time"
)

func asPoint(record *heartbeat.HeartBeat, id string, ts time.Time) *write.Point {
	return write.NewPoint(
		"cow", map[string]string{"uuid": id}, map[string]interface{}{
			"weight":    record.GetWeight(),
			"health":    record.GetHealth(),
			"latitude":  record.GetLatitude(),
			"longitude": record.GetLongitude(),
		}, ts,
	)
}

func HeartBeatHandler(server mqtt.Client, msg mqtt.Message) {
	heartBeat := new(heartbeat.HeartBeat)
	if err := proto.Unmarshal(msg.Payload(), heartBeat); err != nil {
		Logger.Warnw("heartbeat error: unmarshal error", "error", err.Error())
		return
	}
	claims, err := utils.JWTParse(heartBeat.GetToken())
	if err != nil {
		Logger.Warnw("heartbeat error: parse token error", "error", err.Error())
		return
	}
	ts, err := time.Parse(time.RFC3339, heartBeat.GetTimestamp())
	if err != nil {
		Logger.Warnw("heartbeat error: parse timestamp error", "error", err.Error())
		return
	}
	point := asPoint(heartBeat, claims.Id, ts)
	db.InfluxWriteApi.WritePoint(point)
	stream.Input() <- heartBeat
	if !constant.InBound(heartBeat.Longitude, heartBeat.Latitude) {
		topic := fmt.Sprintf("cow/%s/command/banish", claims.Id)
		server.Publish(topic, 0, false, []byte{})
	}
}
