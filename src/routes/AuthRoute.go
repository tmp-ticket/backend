package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/tmp-ticket/backend/src/account"
	"github.com/tmp-ticket/backend/src/datastructs"
)

func AuthUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			slog.Info(err.Error())
			return
		}

		var accountInfo datastructs.AccountInfo

		err = json.Unmarshal(body, &accountInfo)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			slog.Info(fmt.Sprintf("Bad request recieved: %s", err.Error()))
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
		} else if verifyAccount.IsAuth() {
			signed, err := account.CreateWebToken(verifyAccount.Id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				slog.Error(err.Error())
				return
			}
			http.Redirect(w, r, "/tickets", http.StatusMovedPermanently)
			cookie := http.Cookie{
				Name:     "Token",
				Value:    *signed,
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteStrictMode,
			}
			http.SetCookie(w, &cookie)
			slog.Info(fmt.Sprintf("User %s successfully authenticated", accountInfo.Email))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
