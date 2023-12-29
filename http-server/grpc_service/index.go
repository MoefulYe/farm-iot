package grpc_service

import (
	"github.com/MoefulYe/farm-iot/http-server/config"
	"github.com/MoefulYe/farm-iot/http-server/logger"
	"github.com/MoefulYe/farm-iot/http-server/protoc-gen/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Client service.ServiceClient
	conn   *grpc.ClientConn
)

func init() {
	_conn, err := grpc.Dial(config.Conf.GrpcServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Fatalw(err.Error())
	}
	conn = _conn
	Client = service.NewServiceClient(conn)
}
