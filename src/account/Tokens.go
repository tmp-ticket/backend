package account

import (
	"errors"
	"log/slog"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/ed25519"
)

var priv ed25519.PrivateKey = nil
var pub ed25519.PublicKey = nil

func CreateWebToken(Id int) (*string, error) {

	if priv == nil || pub == nil {
		err := setupKeysRand()
		if err != nil {
			return nil, err
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{"sub": Id})
	signed, err := token.SignedString(&priv)

	if err != nil {
		return nil, err
	}

	return &signed, nil

}

func setupKeysRand() error {
	if priv == nil && pub == nil {
		var err error
		pub, priv, err = ed25519.GenerateKey(nil)
		if err != nil {
			slog.Error(err.Error())
			return err
		}
	}
	slog.Error("cannot generate signature keys")
	return errors.New("cannot generate signature keys")
}
