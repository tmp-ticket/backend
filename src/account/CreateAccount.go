package account

import (
	"github.com/tmp-ticket/backend/src/database"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(email string, password string) error {

	hashpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	database.CreateAccountDB(email, hashpass)
	return nil
}
