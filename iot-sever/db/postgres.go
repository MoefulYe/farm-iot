package db

import (
	"context"
	"github.com/MoefulYe/farm-iot/database/postgres/ent"
	"github.com/MoefulYe/farm-iot/iot-server/config"
	_ "github.com/lib/pq"
	"log"
)

var (
	PgClient *ent.Client
)

func init() {
	client, err := ent.Open("postgres", config.Conf.Postgres)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
	PgClient = client
}
