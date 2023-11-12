package db

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var InfluxClient influxdb2.Client
var InfluxWriteApi api.WriteAPI
var QueryAPI api.QueryAPI

func init() {
	InfluxClient = influxdb2.NewClient("http://124.221.89.92:8086", "EFr0bgEQip96t-RL99r6rURvBzj0MFi4LtC-vCpIKaQYu4CjKm5M59xXakfL2NtLMArwPlXUykhrinwJVD53Zg==")
	InfluxWriteApi = InfluxClient.WriteAPI("farm-iot", "farm-iot")
	QueryAPI = InfluxClient.QueryAPI("farm-iot")
}
