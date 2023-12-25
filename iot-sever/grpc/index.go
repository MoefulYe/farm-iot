package grpc

import (
	"fmt"
	"github.com/MoefulYe/farm-iot/iot-server/config"
	"github.com/MoefulYe/farm-iot/iot-server/logger"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/grpc/service"
	"github.com/MoefulYe/farm-iot/iot-server/server"
	"google.golang.org/grpc"
	"net"
)
import "context"

var listener net.Listener
var serv *grpc.Server

func init() {
	_listener, err := net.Listen("tcp", config.Conf.GrpcAddr)
	if err != nil {
		logger.Logger.Fatalw(err.Error())
	}
	listener = _listener
	serv = grpc.NewServer()
	service.RegisterServiceServer(serv, &grpcService{})
	go func() {
		if err = serv.Serve(listener); err != nil {
			logger.Logger.Error(err.Error())
		}
	}()
}

type grpcService struct {
	service.UnimplementedServiceServer
}

func (*grpcService) Spawn(_ context.Context, _ *service.SpawnReq) (*service.SpawnResp, error) {
	if token := server.Server.Publish("spawner/spawn", 0, false, []byte{}); token.Wait() && token.Error() != nil {
		err := token.Error()
		logger.Logger.Warnw(err.Error())
		return nil, err
	} else {
		return &service.SpawnResp{}, nil
	}
}

func (*grpcService) Kill(_ context.Context, req *service.KillReq) (*service.KillResp, error) {
	for _, id := range req.List {
		topic := fmt.Sprintf("cow/%s/command/kill", id)
		if token := server.Server.Publish(topic, 0, false, []byte{}); token.Wait() && token.Error() != nil {
			err := token.Error()
			logger.Logger.Warnw(err.Error())
			return nil, err
		}
	}
	return &service.KillResp{}, nil
}
