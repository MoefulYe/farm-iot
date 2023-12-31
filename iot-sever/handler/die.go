package handler

import (
	. "github.com/MoefulYe/farm-iot/iot-server/ctx"
	"github.com/MoefulYe/farm-iot/iot-server/db"
	"github.com/MoefulYe/farm-iot/iot-server/logger"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow/command"
	"github.com/MoefulYe/farm-iot/iot-server/stream"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
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
		stream.IncomeStream() <- money
	}
}
