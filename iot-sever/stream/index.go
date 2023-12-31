package stream

import (
	"github.com/MoefulYe/farm-iot/iot-server/db"
	"github.com/MoefulYe/farm-iot/iot-server/logger"
	"github.com/MoefulYe/farm-iot/iot-server/protoc-gen/farm/cow/heartbeat"
	"github.com/MoefulYe/farm-iot/iot-server/server"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	ext "github.com/reugn/go-streams/extension"
	"github.com/reugn/go-streams/flow"
	"time"
)

var (
	source = make(chan any, 1024)
	sink   = make(chan any, 1024)
)

func init() {
	go handle()
	go func() {
		ext.NewChanSource(source).Via(flow.NewTumblingWindow(time.Minute * 5)).To(ext.NewChanSink(sink))
	}()
	logger.Logger.Infow("init data stream")
}

func Input() chan<- any {
	return source
}

func handle() {
	for window := range sink {
		window := window.([]any)
		cnt := len(window)
		sum := 0.0
		if cnt > 0 {
			feedOutcome := 5.0 * cnt
			feed := write.NewPoint(
				"outcome", map[string]string{
					"type": "feed",
				}, map[string]interface{}{
					"out": feedOutcome,
				}, time.Now(),
			)
			db.InfluxWriteApi.WritePoint(feed)
			for _, elem := range window {
				sum += elem.(*heartbeat.HeartBeat).Health
			}
			health := sum / float64(cnt)
			if health < 0.5 {
				logger.Logger.Infow("cure")
				cureOutcome := 50.0 * cnt
				cure := write.NewPoint(
					"outcome", map[string]string{
						"type": "cure",
					}, map[string]interface{}{
						"out": cureOutcome,
					}, time.Now(),
				)
				db.InfluxWriteApi.WritePoint(cure)
				if token := server.Server.Publish(
					"cow/broadcast/command/cure",
					0,
					false,
					[]byte{},
				); token.Wait() && token.Error() != nil {
					logger.Logger.Warnw(token.Error().Error())
				}
			}
		}
	}
}
