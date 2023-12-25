package handler

import (
	"fmt"
	. "github.com/MoefulYe/farm-iot/iot-server/ctx"
	"github.com/MoefulYe/farm-iot/iot-server/db"
	"github.com/MoefulYe/farm-iot/iot-server/logger"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow/command"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"time"
)

func DieHandler(_ mqtt.Client, msg mqtt.Message) {
	die := new(command.Die)
	if err := proto.Unmarshal(msg.Payload(), die); err != nil {
		logger.Logger.Warnw(err.Error())
	}
	id := uuid.MustParse(die.Uuid)
	ts, err := time.Parse(time.RFC3339, die.Timestamp)
	if err != nil {
		logger.Logger.Warnw(err.Error())
	}
	if err = db.PgClient.Device.UpdateOneID(id).SetDeadAt(ts).SetReason(die.Reason).Exec(Ctx); err != nil {
		logger.Logger.Warnw(err.Error())
	}

	if die.Reason == "kill" {
		money := die.Weight * 30.0
		reason := fmt.Sprintf("kill cow-%s and get %v", die.Uuid, money)
		point := write.NewPoint(
			"balance", map[string]string{}, map[string]interface{}{
				"reason":  reason,
				"balance": money,
			}, ts,
		)
		db.InfluxWriteApi.WritePoint(point)
	}
}
