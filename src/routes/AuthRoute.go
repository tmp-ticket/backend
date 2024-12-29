package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/tmp-ticket/backend/src/account"
)

type AccountInfo struct {
	Email    string
	Password string
}

func AuthUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			slog.Info(err.Error())
			return
		}

		var accountInfo AccountInfo

		err = json.Unmarshal(body, &accountInfo)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			slog.Info(err.Error())
			return
		}
		verifyAccount, err := account.AuthAccount(accountInfo.Email, accountInfo.Password)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			slog.Error(err.Error())
			return
		}

		if !verifyAccount.IsAuth() {
			w.WriteHeader(http.StatusForbidden)
			slog.Info(fmt.Sprintf("User %s failed to authenticate", accountInfo.Email))
			return
		}

		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		slog.Info(fmt.Sprintf("User %s successfully authenticated", accountInfo.Email))
		return
	} else {
		return
	}
}
