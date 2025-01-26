package database

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGetDBAccountEmail(t *testing.T) {

	resetDatabase()
	hashpass, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	err := CreateAccountDB("test@test.com", hashpass)

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	data, err := GetDBAccountEmail("test@test.com")

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if data.Email != "test@test.com" {
		t.Errorf("Email is incorrect")
	}

	if bcrypt.CompareHashAndPassword(data.Password, []byte("test")) != nil {
		t.Errorf("password is incorrect")
	}
}
