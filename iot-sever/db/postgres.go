package db

import (
	"github.com/MoefulYe/farm-iot/database/postgres/ent"
	"github.com/MoefulYe/farm-iot/iot-server/config"
	"log"
)

var (
	PgClient *ent.Client
)

func init() {
	client, err := ent.Open("postgres", config.Conf.PgConnStr)
	if err != nil {
		log.Fatalf(err.Error())
	}
	PgClient = client
}
