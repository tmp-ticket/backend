package database

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	err := SetupDBConnection()

	if err != nil {
		t.Errorf("recieved error on connection: %s", err.Error())
	}

	if DB_POOL == nil {
		t.Errorf("the database pool is nil after calling setup function")
	}
}
