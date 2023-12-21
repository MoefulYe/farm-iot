package db

import (
	"context"
	"github.com/MoefulYe/farm-iot/database/postgres/ent"
	"github.com/MoefulYe/farm-iot/iot-server/config"
	. "github.com/MoefulYe/farm-iot/iot-server/logger"
	_ "github.com/lib/pq"
)

var (
	PgClient *ent.Client
)

func init() {
	client, err := ent.Open("postgres", config.Conf.Postgres)
	if err != nil {
		Logger.Fatal(err.Error())
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		Logger.Fatal(err.Error())
	}
	PgClient = client
}
