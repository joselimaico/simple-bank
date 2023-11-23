package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store
const dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

func TestMain(m *testing.M) {
	var err error
	ctx := context.Background()

	connPool, err := pgxpool.New(ctx, dbSource)

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}