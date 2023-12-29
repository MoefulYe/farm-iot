package db

import (
	"github.com/MoefulYe/farm-iot/http-server/config"
	"github.com/MoefulYe/farm-iot/http-server/logger"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var Client influxdb2.Client
var WriteApi api.WriteAPI
var QueryAPI api.QueryAPI

func init() {
	Client = influxdb2.NewClient(
		config.Conf.Influxdb.Url,
		config.Conf.Influxdb.Auth,
	)
	WriteApi = Client.WriteAPI("farm-iot", "farm")
	QueryAPI = Client.QueryAPI("farm-iot")
	go func() {
		for err := range WriteApi.Errors() {
			logger.Logger.Errorw(err.Error())
		}
	}()
	logger.Logger.Infow("influxdb connected")
}
