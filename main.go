package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joselimaico/simplebank/api"
	db "github.com/joselimaico/simplebank/db/sqlc"
	"github.com/joselimaico/simplebank/util"
)



func main() {
	var err error
	ctx := context.Background()
	config,err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	connPool, err := pgxpool.New(ctx, config.DBSource)

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server: ", err)
	}

}
