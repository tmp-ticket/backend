package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tmp-ticket/backend/src/account"
)

type AccountInfo struct {
	email    string
	password string
}

func AuthUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var accountInfo AccountInfo

		err = json.Unmarshal(body, accountInfo)
		verifyAccount, err := account.AuthAccount(accountInfo.email, accountInfo.password)

		if !verifyAccount.IsAuth() {
			w.WriteHeader(http.StatusForbidden)
		}

		http.Redirect(w, r, "/home", http.StatusMovedPermanently)

		return
	} else {
		return
	}
}
