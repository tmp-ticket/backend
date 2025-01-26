package database

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/tmp-ticket/backend/src/datastructs"
)

func GetDBAccount(id int) (*datastructs.RawAccount, error) {
	if DB_POOL == nil {
		err := SetupDBConnection()
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}
	}

	conn, err := DB_POOL.GetConn()

	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("database has failed to connect while gettting account by email")
	}
	defer conn.Release()
	data, err := conn.Query(context.Background(), "SELECT * FROM accounts WHERE id = $1", id)

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	account, err := pgx.CollectExactlyOneRow(data, pgx.RowToAddrOfStructByName[datastructs.RawAccount])

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return account, nil

}

func GetDBAccountEmail(email string) (*datastructs.RawAccount, error) {

	if DB_POOL == nil {
		err := SetupDBConnection()
		if err != nil {
			slog.Error(err.Error())
			return nil, err
		}
	}

	conn, err := DB_POOL.GetConn()

	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("database has failed to connect while gettting account by email")
	}
	defer conn.Release()
	data, err := conn.Query(context.Background(), "SELECT * FROM accounts WHERE email = $1", email)

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	account, err := pgx.CollectExactlyOneRow(data, pgx.RowToAddrOfStructByName[datastructs.RawAccount])

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return account, nil
}
