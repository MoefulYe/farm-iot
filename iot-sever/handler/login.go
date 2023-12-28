package handler

import (
	"fmt"
	. "github.com/MoefulYe/farm-iot/iot-server/ctx"
	"github.com/MoefulYe/farm-iot/iot-server/db"
	. "github.com/MoefulYe/farm-iot/iot-server/logger"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow/login"
	"github.com/MoefulYe/farm-iot/iot-server/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"pg/ent/device"
)

func LoginHandler(server mqtt.Client, msg mqtt.Message) {
	payload := msg.Payload()
	req := new(login.LoginReq)
	if err := proto.Unmarshal(payload, req); err != nil {
		Logger.Warnw("login error: fail to unmarshal payload", "error", err.Error())
		return
	}
	id := uuid.MustParse(req.GetUuid())
	passwd := req.GetPasswd()
	d, err := db.PgClient.Device.Query().Where(device.IDEQ(id)).First(Ctx)
	if err != nil {
		Logger.Warnw("login error: fail to query device", "error", err.Error())
		handleLoginError(server, req.GetUuid(), login.Status_STATUS_INVALID_UUID)
		return
	}
	if match, err := utils.Argon2Verify(passwd, d.HashedPasswd); match == false || err != nil {
		Logger.Warnw("login error: fail to password verify", "error", err.Error())
		handleLoginError(server, req.GetUuid(), login.Status_STATUS_INVALID_PASSWD)
		return
	}
	token, err := utils.JWTGenerate(utils.NewClaims(req.GetUuid()))
	if err != nil {
		Logger.Warnw("login error: fail to generate token", "error", err.Error())
		handleLoginError(server, req.GetUuid(), login.Status_STATUS_INVALID_PASSWD)
		return
	}
	resp := &login.LoginResp{
		Status: login.Status_STATUS_OK,
		Token:  token,
	}
	toSend, err := proto.Marshal(resp)
	if err != nil {
		Logger.Warnw("login error: fail to marshal response", "error", err.Error())
		handleLoginError(server, req.GetUuid(), login.Status_STATUS_INVALID_PASSWD)
		return
	}
	topic := fmt.Sprintf("cow/%s/login-reply", req.GetUuid())
	if tok := server.Publish(topic, 0, false, toSend); tok.Wait() && tok.Error() != nil {
		Logger.Warnw("login error: fail to publish response", "error", tok.Error().Error())
		return
	}

}

func handleLoginError(server mqtt.Client, id string, status login.Status) {
	topic := fmt.Sprintf("cow/%s/login-reply", id)
	resp := &login.LoginResp{
		Status: status,
		Token:  "",
	}
	payload, err := proto.Marshal(resp)
	if err != nil {
		Logger.Warnw("login error: fail to marshal response", "error", err.Error())
		return
	}
	if token := server.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		Logger.Warnw("login error: fail to publish response", "error", token.Error().Error())
		return
	}
}
