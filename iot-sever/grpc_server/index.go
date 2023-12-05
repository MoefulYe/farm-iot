package grpc_server

import (
	"context"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/grpc/kill"
	mqtt "github.com/MoefulYe/farm-iot/iot-server/server"
	"github.com/golang/protobuf/proto"
)

type Server struct {
	kill.UnimplementedKillServer
}

func (s *Server) Kill(c context.Context, req *kill.KillReq) (*kill.KillResp, error) {
	toSend, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	if token := mqtt.Server.Publish("killer/kill", 0, false, toSend); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &kill.KillResp{}, nil
}
