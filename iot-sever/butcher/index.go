package butcher

import (
	"pg/ent"
	"pg/ent/device"
	"time"

	"github.com/MoefulYe/farm-iot/iot-server/ctx"
	"github.com/MoefulYe/farm-iot/iot-server/db"
	"github.com/MoefulYe/farm-iot/iot-server/logger"
)

const duration = time.Minute * 30

func init() {
	go func() {
		ticker := time.NewTicker(duration)
		butcher()
		for range ticker.C {
			butcher()
		}
	}()
}

func butcher() {
	cnt, err := db.PgClient.Device.Query().Where(
		device.DeadAtIsNil(),
	).Aggregate(
		ent.Count(),
	).Int(ctx.Ctx)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		return
	}

	if cnt < 100 {
		logger.Logger.Infof("cow cnt is %v and less than 100. no cow is killed", cnt)
	} else if cnt <= 200 {

	} else if cnt <= 300 {

	} else if cnt <= 400 {

	}
}
