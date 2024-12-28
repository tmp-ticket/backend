package account

type RawAccount struct {
	email    string
	id       int
	password string
}

func GetAccount(id int) (*RawAccount, error) { return nil, nil }

func GetAccountEmail(email string) (*RawAccount, error) { return nil, nil }
