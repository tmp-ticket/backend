package account

import (
	"github.com/tmp-ticket/backend/src/database"
	"github.com/tmp-ticket/backend/src/datastructs"
)

func GetAccount(id int) (*datastructs.RawAccount, error) {

	account, err := database.GetDBAccount(id)

	return account, err
}

func GetAccountEmail(email string) (*datastructs.RawAccount, error) {

	account, err := database.GetDBAccountEmail(email)

	return account, err
}
