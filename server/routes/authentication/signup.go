package authentication

import (
	"cdm/server/models"
	"cdm/server/muxrouter"
	"net/http"
	"log"
)

func fasterErrors(err error) muxrouter.JSON {
	return muxrouter.JSON{"errors": muxrouter.JSON{"global": err.Error()}}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	// TODO email confirmation send u.SendEmailConfirmation()
	var c models.UserCredentials
	ww := muxrouter.Mhrw{ResponseWriter: w}

	if err := c.GetCredentials(r.Body); err != nil {
		ww.Error(fasterErrors(err))
		return
	}
	if err := c.Credentials.InsertNewUser(); err != nil {
		ww.Error(fasterErrors(err))
		return
	}
	log.Println(c.Credentials.GenerateConfirmationURL())
	ww.Success(muxrouter.JSON{"user": c.Credentials.ToAuthJSON()})
}
