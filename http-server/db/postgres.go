package db

import (
	"context"
	"fmt"
	"github.com/MoefulYe/farm-iot/database/postgres/ent"
	"github.com/MoefulYe/farm-iot/http-server/config"
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
	fmt.Printf("ent ok")
	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
	PgClient = client
}
