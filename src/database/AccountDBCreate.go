package database

import (
	"context"
	"errors"
	"log/slog"
)

func CreateAccountDB(email string, password []byte) error {

	if DB_POOL == nil {
		SetupDBConnection()
	}

	conn, err := DB_POOL.GetConn()

	defer conn.Release()
	if err != nil {
		slog.Error(err.Error())

		return errors.New("database connection has failed")
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		slog.Error(err.Error())

		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "INSERT INTO accounts VALUES (DEFAULT, $1, $2)", email, password)
	if err != nil {
		slog.Error(err.Error())

		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		slog.Error(err.Error())

		return err
	}

	return nil
}
