package authentication

import (
	"net/http"
	"cdm/server/muxrouter"
	"cdm/server/models"
)

func ValidateToken (w http.ResponseWriter, r *http.Request) {
	ww := muxrouter.Mhrw{ResponseWriter: w}
	token, err := getToken(r.Body)
	if err != nil {
		ww.Error(fasterErrors(err))
		return
	}
	// TODO compare with db token
	u :=models.User{ConfirmationToken:token}
	if err := u.ValidateToken(); err != nil {
		ww.Error401(nil)
		return
	}
	ww.Success(nil)
}
