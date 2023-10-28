package handler

import (
	"context"
	"fmt"
	"github.com/MoefulYe/farm-iot/iot-server/db"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow"
	"github.com/MoefulYe/farm-iot/iot-server/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"time"
)

func RegisterHandler(server mqtt.Client, msg mqtt.Message) {
	payload := new(cow.RegisterReq)
	err := proto.Unmarshal(msg.Payload(), payload)
	if err != nil {
		return
	}
	passwd := payload.GetPasswd()
	bornAt, err := time.Parse(time.RFC3339, payload.GetBornAt())
	if err != nil {
		return
	}
	id, err := uuid.Parse(payload.GetUuid())
	if err != nil {
		return
	}
	hashedPasswd, err := utils.Argon2Generate(passwd)
	if err != nil {
		handleRegisterResult(
			server, payload.GetUuid(), cow.RegisterResp_FAILED, "",
		)
		return
	}
	if _, err = db.PgClient.Device.Create().SetID(id).SetBornAt(bornAt).SetHashedPasswd(hashedPasswd).Save(context.Background()); err != nil {
		handleRegisterResult(
			server, payload.GetUuid(), cow.RegisterResp_FAILED, "",
		)
		return
	}
	claims := utils.NewClaims(payload.GetUuid())
	token, err := utils.JWTGenerate(claims)
	if err != nil {
		handleRegisterResult(
			server, payload.GetUuid(), cow.RegisterResp_FAILED, "",
		)
		return
	}
	if err != nil {
		handleRegisterResult(
			server, payload.GetUuid(), cow.RegisterResp_FAILED, "",
		)
		return
	}
	handleRegisterResult(server, payload.GetUuid(), cow.RegisterResp_OK, token)
}

func handleRegisterResult(
	server mqtt.Client, id string, status cow.RegisterResp_Status, token string,
) {
	topic := fmt.Sprintf("cow/%s/register-reply", id)
	resp := &cow.RegisterResp{
		Status: status,
		Uuid:   id,
		Token:  token,
	}
	payload, err := proto.Marshal(resp)
	if err != nil {
		return
	}
	if tok := server.Publish(
		topic, 0, false, payload,
	); tok.Wait() && tok.Error() != nil {
		return
	}
}
