package datastructs

import (
	"errors"
	"net/mail"
)

type Account struct {
	Email   string
	Id      int
	Is_auth bool
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
	return &Account{Email: email, Id: id, Is_auth: auth}, nil
}

func (account Account) IsAuth() bool {
	return account.Is_auth
}

func (account Account) GetID() int {
	return account.Id
}

func (account Account) GetEmail() string {
	return account.Email
}
