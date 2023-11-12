package db

import (
	"context"
	"fmt"
	"github.com/MoefulYe/farm-iot/database/postgres/ent"
	_ "github.com/lib/pq"
	"log"
)

var (
	PgClient *ent.Client
)

func init() {
	client, err := ent.Open("postgres", "host=124.221.89.92 port=5432 user=farmer password=mysecretpassword dbname=farm-iot sslmode=disable")
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
