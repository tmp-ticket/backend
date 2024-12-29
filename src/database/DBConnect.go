package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBManager struct {
	pool *pgxpool.Pool
}

var DB_POOL *DBManager

func SetupDBConnection() error {

	pool, err := pgxpool.New(context.Background(), fmt.Sprintf("user=%s password=%s host=%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME")))

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	test_conn, err := pool.Acquire(context.Background())

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = test_conn.Ping(context.Background())

	if err != nil {

		slog.Error(err.Error())
		os.Exit(1)
	}

	test_conn.Release()
	test_conn = nil

	DB_POOL = &DBManager{pool: pool}

	return err
}

func (db DBManager) GetConn() (*pgxpool.Conn, error) {
	conn, err := db.pool.Acquire(context.Background())
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return conn, err
}

func (db DBManager) CloseDBPool() {
	db.pool.Close()
}
