package db

import (
	"github.com/MoefulYe/farm-iot/http-server/config"
	"github.com/MoefulYe/farm-iot/http-server/ctx"
	"github.com/MoefulYe/farm-iot/http-server/logger"
	_ "github.com/lib/pq"
	"pg/ent"
)

var (
	PgClient *ent.Client
)

func init() {
	client, err := ent.Open(
		"postgres",
		config.Conf.Postgres,
	)
	if err != nil {
		logger.Logger.Fatalw(err.Error())
	}
	err = client.Schema.Create(ctx.Bg)
	if err != nil {
		logger.Logger.Fatalw(err.Error())
	}
	PgClient = client
	logger.Logger.Infow("postgres connected")
}
