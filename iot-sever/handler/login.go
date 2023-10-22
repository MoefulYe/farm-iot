package handler

import (
	"context"
	"fmt"
	"github.com/MoefulYe/farm-iot/database/postgres/ent/device"
	"github.com/MoefulYe/farm-iot/iot-server/db"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow"
	"github.com/MoefulYe/farm-iot/iot-server/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

func LoginHandler(server mqtt.Client, msg mqtt.Message) {
	payload := msg.Payload()
	req := new(cow.LoginReq)
	if err := proto.Unmarshal(payload, req); err != nil {
		return
	}
	id := uuid.MustParse(req.GetUuid())
	passwd := req.GetPasswd()
	d, err := db.PgClient.Device.Query().Where(device.IDEQ(id)).First(context.Background())
	if err != nil {
		handleLoginError(server, req.GetUuid(), cow.LoginResp_INVALID_UUID)
		return
	}
	if match, err := utils.Argon2Verify(passwd, d.HashedPasswd); match == false || err != nil {
		handleLoginError(server, req.GetUuid(), cow.LoginResp_INVALID_PASSWD)
		return
	}
	token, err := utils.JWTGenerate(utils.NewClaims(req.GetUuid()))
	if err != nil {
		handleLoginError(server, req.GetUuid(), cow.LoginResp_INVALID_PASSWD)
		return
	}
	resp := &cow.LoginResp{
		Status: cow.LoginResp_OK,
		Token:  token,
	}
	toSend, err := proto.Marshal(resp)
	if err != nil {
		handleLoginError(server, req.GetUuid(), cow.LoginResp_INVALID_PASSWD)
		return
	}
	topic := fmt.Sprintf("cow/%s/login-reply", req.GetUuid())
	if tok := server.Publish(topic, 0, false, toSend); tok.Wait() && tok.Error() != nil {
		return
	}

}

func handleLoginError(server mqtt.Client, id string, status cow.LoginResp_Status) {
	topic := fmt.Sprintf("cow/%s/login-reply", id)
	resp := &cow.LoginResp{
		Status: status,
		Token:  "",
	}
	payload, err := proto.Marshal(resp)
	if err != nil {
		return
	}
	if token := server.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		return
	}
}
