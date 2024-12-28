package account

import (
	"errors"
	"net/mail"
)

type Account struct {
	email   string
	id      int
	is_auth bool
}

func NewAccountStruct(email string, id int, auth bool) (*Account, error) {
	if email == "" {
		return nil, errors.New("no email address provided")
	}
	_, mailCheck := mail.ParseAddress(email)

	if mailCheck != nil {
		return nil, errors.New("email provided is invalid")
	}

	if id < 1 {
		return nil, errors.New("id is invalid")
	}
	return &Account{email: email, id: id, is_auth: auth}, nil
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
