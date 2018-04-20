package authentication

import (
	"cdm/server/models"
	"cdm/server/muxrouter"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type datas struct {
	A a `json:"data"`
}
type a struct {
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	Token                string `json:"token"`
}

func getPassword(reader io.Reader) (a, error) {
	var d datas
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return a{}, err
	}
	if err := json.Unmarshal(body, &d); err != nil {
		return a{}, err
	}
	log.Println(d)
	return d.A, nil
}

/*
ResetPassword function
*/
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	pw, err := getPassword(r.Body)
	ww := muxrouter.Mhrw{ResponseWriter: w}
	if err != nil {
		ww.Error(fasterErrors(err))
		return
	}
	if pw.Password == "" || pw.PasswordConfirmation == "" || pw.Token == "" {
		ww.Error(fasterErrors(errors.New("aie")))
	}
	u := models.User{
		ConfirmationToken: pw.Token,
	}
	if err := u.ValidateToken(); err != nil {
		ww.Error(fasterErrors(err))
		return
	}
	if err := u.FindWithId(pw.Token); err != nil {
		ww.Error(fasterErrors(err))
		return
	}
	// TODO compare old password with new if it is the same
	u.PasswordHash = pw.Password
	u.SetPassword()
	if err := u.Update(); err != nil {
		ww.Error(fasterErrors(err))
		return

	}
	ww.Success(nil)
}
