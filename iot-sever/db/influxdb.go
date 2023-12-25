package db

import (
	"github.com/MoefulYe/farm-iot/iot-server/config"
	"github.com/MoefulYe/farm-iot/iot-server/logger"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var InfluxClient influxdb2.Client
var InfluxWriteApi api.WriteAPI

func init() {
	opts := config.Conf.Influxdb
	InfluxClient = influxdb2.NewClient(opts.Url, opts.Auth)
	InfluxWriteApi = InfluxClient.WriteAPI("farm-iot", "farm")
	go func() {
		for err := range InfluxWriteApi.Errors() {
			logger.Logger.Warnw(err.Error())
		}
	}()
	logger.Logger.Infow("init influxdb connection")
}
