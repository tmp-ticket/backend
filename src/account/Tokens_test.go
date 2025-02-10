package account

import (
	"encoding/base64"
	"strconv"
	"testing"
)

func VerifyTest(t *testing.T) {

	token, err := CreateWebToken(1)

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	parsed_token, err := VerifyToken(*token)

	if err != nil {
		t.Errorf("%s", err.Error())
	}
	b64str, err := parsed_token.Claims.GetSubject()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	b64int, err := base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	id, err := strconv.Atoi(string(b64int))
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if id != 1 {
		t.Errorf("parsed token subject is incorrecct")
	}

}
