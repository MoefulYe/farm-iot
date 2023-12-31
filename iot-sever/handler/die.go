package handler

import (
	. "github.com/MoefulYe/farm-iot/iot-server/ctx"
	"github.com/MoefulYe/farm-iot/iot-server/db"
	"github.com/MoefulYe/farm-iot/iot-server/logger"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow/command"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"google.golang.org/protobuf/proto"
	"time"
)

func DieHandler(_ mqtt.Client, msg mqtt.Message) {
	die := new(command.Die)
	if err := proto.Unmarshal(msg.Payload(), die); err != nil {
		logger.Logger.Warnw(err.Error())
		return
	}
	id := uuid.MustParse(die.Uuid)
	ts, err := time.Parse(time.RFC3339, die.Timestamp)
	if err != nil {
		logger.Logger.Warnw(err.Error())
		return
	}
	if err = db.PgClient.Device.UpdateOneID(id).SetDeadAt(ts).SetReason(die.Reason).Exec(Ctx); err != nil {
		logger.Logger.Warnw(err.Error())
		return
	}

	if die.Reason == "kill" {
		money := die.Weight * 30.0
		point := write.NewPoint(
			"income", map[string]string{
				"type": "kill",
			}, map[string]interface{}{
				"in": money,
			}, ts,
		)
		db.InfluxWriteApi.WritePoint(point)
	}
}
