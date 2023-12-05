package handler

import (
	"context"
	"github.com/MoefulYe/farm-iot/iot-server/db"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/die"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	uuid2 "github.com/google/uuid"
	"time"
)

func DieMsgHandler(server mqtt.Client, msg mqtt.Message) {
	req := new(die.DieMsg)
	if err := proto.Unmarshal(msg.Payload(), req); err != nil {
		return
	}
	reason := req.GetRecord().Reason
	uuid := uuid2.MustParse(req.GetRecord().Uuid)
	when, err := time.Parse(time.RFC3339, req.GetRecord().Timestamp)
	if err != nil {
		return
	}
	if err = db.PgClient.Device.UpdateOneID(uuid).SetDeadAt(when).SetReason(reason).Exec(context.Background()); err != nil {
		return
	}
}
