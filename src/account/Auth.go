package account

func AuthAccount(email string, password string) (*Account, error) {
	account, err := GetAccountEmail(email)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, nil
	}
	// check if account password is correct

	return nil, nil
}
