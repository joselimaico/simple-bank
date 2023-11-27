package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joselimaico/simplebank/util"
)

var testStore Store

func TestMain(m *testing.M) {
	var err error
	ctx := context.Background()
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config ",err)
	}
	connPool, err := pgxpool.New(ctx, config.DBSource)

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}