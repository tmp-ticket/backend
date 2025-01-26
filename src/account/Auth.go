package account

import (
	"github.com/tmp-ticket/backend/src/datastructs"
	"golang.org/x/crypto/bcrypt"
)

func AuthAccount(email string, password string) (*datastructs.Account, error) {
	check_acc, err := GetAccountEmail(email)
	if err != nil {
		return nil, err
	}

	if check_acc == nil {
		return nil, nil
	}
	// check if account password is correct

	err = bcrypt.CompareHashAndPassword(check_acc.Password, []byte(password))

	if err != nil {
		return nil, err
	}

	ret_acc := &datastructs.Account{
		Email:   email,
		Id:      check_acc.Id,
		Is_auth: true,
	}

	return ret_acc, nil
}
