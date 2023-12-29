package handler

import (
	"fmt"
	. "github.com/MoefulYe/farm-iot/iot-server/ctx"
	"github.com/MoefulYe/farm-iot/iot-server/db"
	. "github.com/MoefulYe/farm-iot/iot-server/logger"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow/register"
	"github.com/MoefulYe/farm-iot/iot-server/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"time"
)

func RegisterHandler(server mqtt.Client, msg mqtt.Message) {
	payload := new(register.RegisterReq)
	err := proto.Unmarshal(msg.Payload(), payload)
	if err != nil {
		Logger.Warnw("register error: unmarshal error", "error", err.Error())
		return
	}
	passwd := payload.GetPasswd()
	bornAt, err := time.Parse(time.RFC3339, payload.GetBornAt())
	if err != nil {
		Logger.Warnw("register error: parse born_at error", "error", err.Error())
		return
	}
	id, err := uuid.Parse(payload.GetUuid())
	if err != nil {
		Logger.Warnw("register error: parse uuid error", "error", err.Error())
		return
	}
	hashedPasswd, err := utils.Argon2Generate(passwd)
	if err != nil {
		Logger.Warnw("register error: generate hashed passwd error", "error", err.Error())
		handleRegisterResult(
			server, payload.GetUuid(), register.Status_STATUS_FAILED, "",
		)
		return
	}
	parent := payload.GetParent()
	query := db.PgClient.Device.Create().SetID(id).SetBornAt(bornAt).SetHashedPasswd(hashedPasswd)
	if parent == "" {
		if _, err = query.Save(Ctx); err != nil {
			Logger.Warnw("register error: save device error", "error", err.Error())
			handleRegisterResult(server, payload.GetUuid(), register.Status_STATUS_FAILED, "")
			return
		}
	} else {
		parent, err := uuid.Parse(parent)
		if err != nil {
			Logger.Warnw("register error: parse parent error", "error", err.Error())
			handleRegisterResult(server, payload.GetUuid(), register.Status_STATUS_FAILED, "")
			return
		}
		if _, err = query.SetMotherID(parent).Save(Ctx); err != nil {
			Logger.Warnw("register error: save device error", "error", err.Error())
			handleRegisterResult(server, payload.GetUuid(), register.Status_STATUS_FAILED, "")
			return
		}
	}
	claims := utils.NewClaims(payload.GetUuid())
	token, err := utils.JWTGenerate(claims)
	if err != nil {
		Logger.Warnw("register error: fail to gen token")
		handleRegisterResult(
			server, payload.GetUuid(), register.Status_STATUS_FAILED, "",
		)
		return
	}
	handleRegisterResult(server, payload.GetUuid(), register.Status_STATUS_OK, token)
}

func handleRegisterResult(
	server mqtt.Client, id string, status register.Status, token string,
) {
	topic := fmt.Sprintf("cow/%s/register-reply", id)
	resp := &register.RegisterResp{
		Status: status,
		Uuid:   id,
		Token:  token,
	}
	payload, err := proto.Marshal(resp)
	if err != nil {
		Logger.Warnw("register error: marshal error", "error", err.Error())
		return
	}
	if tok := server.Publish(
		topic, 0, false, payload,
	); tok.Wait() && tok.Error() != nil {
		Logger.Warnw("register error: publish error", "error", tok.Error().Error())
		return
	}
}
