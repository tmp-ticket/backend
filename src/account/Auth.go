package account

func AuthAccount(email string, password string) {
	account, err := GetAccountEmail(email)
	if err != nil {
		return err
	}

	// check if account password is correct
}
