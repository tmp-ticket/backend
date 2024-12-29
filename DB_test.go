package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestDBConnection(t *testing.T) {
	pool, err := pgxpool.New(context.Background(), fmt.Sprintf("user=%s password=%s host=%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME")))

	if err != nil {
		t.Error(err)
	}

	conn, err := pool.Acquire(context.Background())

	if err != nil {
		t.Error(err)
	}

	err = conn.Ping(context.Background())

	if err != nil {
		t.Error(err)
	}
}
