package stream

import (
	"github.com/MoefulYe/farm-iot/iot-server/db"
	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	ext "github.com/reugn/go-streams/extension"
	"github.com/reugn/go-streams/flow"
	"time"
)

var (
	source = make(chan any, 1024)
	sink   = make(chan any, 1024)
)

type KeepAlive struct {
	Uuid      uuid.UUID
	Timestamp time.Time
	Weight    float64
	Health    float64
	Geo       struct {
		Latitude  float64
		Longitude float64
	}
}

func (k *KeepAlive) AsPoint() *write.Point {
	return write.NewPoint(
		"cow", map[string]string{"uuid": k.Uuid.String()},
		map[string]interface{}{
			"weight":    k.Weight,
			"health":    k.Health,
			"latitude":  k.Geo.Latitude,
			"longitude": k.Geo.Longitude,
		}, k.Timestamp,
	)
}

func init() {
	go sendToInfluxdb()
	go func() {
		ext.NewChanSource(source).Via(
			flow.NewMap(
				func(k *KeepAlive) *write.Point {
					return k.AsPoint()
				}, 64,
			),
		).To(ext.NewChanSink(sink))
	}()
}

func Input() chan<- any {
	return source
}

func sendToInfluxdb() {
	for point := range sink {
		db.InfluxWriteApi.WritePoint(point.(*write.Point))
	}
}
