package database

import "context"

func resetDatabase() {

	SetupDBConnection()

	conn, _ := DB_POOL.GetConn()

	conn.Exec(context.Background(), "DROP TABLE accounts")
	conn.Exec(context.Background(), "CREATE TABLE accounts ( ID SERIAL PRIMARY KEY, email varchar(320) UNIQUE, password  bytea)")
}
