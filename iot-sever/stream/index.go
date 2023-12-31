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
	heartbeatSrc  = make(chan any, 1024)
	heartbeatSink = make(chan any, 1024)
	incomeSrc     = make(chan any, 1024)
	incomeSink    = make(chan any, 1024)
)

func init() {
	go handleHeartbeatStream()
	go func() {
		ext.NewChanSource(heartbeatSrc).Via(flow.NewTumblingWindow(time.Minute * 5)).To(ext.NewChanSink(heartbeatSink))
	}()
	go handleIncomeStream()
	go func() {
		ext.NewChanSource(incomeSrc).Via(flow.NewTumblingWindow(time.Minute * 5)).To(ext.NewChanSink(incomeSink))
	}()
	logger.Logger.Infow("init data stream")
}

func HeartbeatStream() chan<- any {
	return heartbeatSrc
}

func IncomeStream() chan<- any { return incomeSrc }

func handleHeartbeatStream() {
	for window := range heartbeatSink {
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

func handleIncomeStream() {
	for window := range incomeSink {
		window := window.([]any)
		sum := 0.0
		for _, income := range window {
			sum += income.(float64)
		}
		point := write.NewPoint(
			"income",
			map[string]string{"type": "kill"},
			map[string]interface{}{"in": sum},
			time.Now(),
		)
		db.InfluxWriteApi.WritePoint(point)
	}
}
