package grpc_service

import (
	"github.com/MoefulYe/farm-iot/http-server/config"
	"github.com/MoefulYe/farm-iot/http-server/protoc-gen/grpc/kill"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	addr   = config.Conf.GrpcServer
	conn   *grpc.ClientConn
	Client kill.KillClient
)

func init() {
	conn_, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(err.Error())
	}
	conn = conn_
	Client = kill.NewKillClient(conn)
}
