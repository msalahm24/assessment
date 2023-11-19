package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"

	"github.com/msalahm24/otp/api"
	db "github.com/msalahm24/otp/db/sqlc"
	"github.com/msalahm24/otp/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not read the config file", err)
	}
	conn, err := pgx.Connect(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer conn.Close(context.Background())

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.HTTPSreverAddress)
	if err != nil {
		log.Fatalf("Can not start the server: %v", err)
	}
}
