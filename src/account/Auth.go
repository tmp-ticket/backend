package account

import "golang.org/x/crypto/bcrypt"

func AuthAccount(email string, password string) (*Account, error) {
	check_acc, err := GetAccountEmail(email)
	if err != nil {
		return nil, err
	}

	if check_acc == nil {
		return nil, nil
	}
	// check if account password is correct

	err = bcrypt.CompareHashAndPassword(check_acc.password, []byte(password))

	if err != nil {
		return nil, err
	}

	ret_acc := &Account{
		email:   email,
		id:      check_acc.id,
		is_auth: true,
	}

	return ret_acc, nil
}
