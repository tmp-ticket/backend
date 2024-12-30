package database

import "testing"

func TestInsertAccount(t *testing.T) {

	err := CreateAccountDB("test@test.com", "test")

	if err != nil {
		t.Errorf("%s", err.Error())
	}
}
