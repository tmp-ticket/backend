package database

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestInsertAccount(t *testing.T) {

	resetDatabase()

	pass, err := bcrypt.GenerateFromPassword([]byte("test"), 0)

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	err = CreateAccountDB("test@test.com", pass)

	if err != nil {
		t.Errorf("%s", err.Error())
	}
}
