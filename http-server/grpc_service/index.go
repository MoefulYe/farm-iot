package grpc_service

import (
	"github.com/MoefulYe/farm-iot/http-server/logger"
	"github.com/MoefulYe/farm-iot/http-server/protoc-gen/grpc/service"
	"github.com/MoefulYe/farm-iot/iot-server/config"
	"google.golang.org/grpc"
)

var (
	Client service.ServiceClient
	conn   *grpc.ClientConn
)

func init() {
	_conn, err := grpc.Dial(config.Conf.GrpcAddr)
	if err != nil {
		logger.Logger.Fatalw(err.Error())
	}
	conn = _conn
	Client = service.NewServiceClient(conn)
}
