package database

import (
	"errors"
	"log/slog"
)

func CreateAccountDB() error {

	_, err := DB_POOL.GetConn()
	if err != nil {
		slog.Error(err.Error())

		return errors.New("Database connection has failed")
	}
	return nil
}
