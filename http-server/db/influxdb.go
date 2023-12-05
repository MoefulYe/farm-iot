package db

import (
	"github.com/MoefulYe/farm-iot/http-server/config"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var InfluxClient influxdb2.Client
var InfluxWriteApi api.WriteAPI
var QueryAPI api.QueryAPI

func init() {
	opts := config.Conf.Influxdb
	InfluxClient = influxdb2.NewClient(opts.Url, opts.Auth)
	InfluxWriteApi = InfluxClient.WriteAPI("farm-iot", "farm-iot")
	QueryAPI = InfluxClient.QueryAPI("farm-iot")
}
