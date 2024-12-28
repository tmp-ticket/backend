package account

type Account struct {
	email   string
	id      int
	is_auth bool
}

func (account Account) IsAuth() bool {
	return account.is_auth
}

func (account Account) GetID() int {
	return account.id
}

func (account Account) GetEmail() string {
	return account.email
}
