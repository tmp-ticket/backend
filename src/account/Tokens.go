package account

import (
	"encoding/base64"
	"errors"
	"log/slog"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	claims := &jwt.RegisteredClaims{
		Subject:   base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(Id))),
		Issuer:    "test", //TODO
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		NotBefore: jwt.NewNumericDate(time.Now().UTC()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signed, err := token.SignedString(&priv)

	if err != nil {
		return nil, err
	}

	return &signed, nil

}

func VerifyToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.ParseWithClaims(token, jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("token not valid")
	}

	return parsedToken, nil
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
