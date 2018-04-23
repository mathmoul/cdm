package muxrouter

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"cdm/server/database/mongoDb"
)

type SU struct {
	NewUser `json:"user"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (su *SU) getNewUser(r io.Reader) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, &su); err != nil {
		return err
	}
	return nil

}

func Signup(w http.ResponseWriter, r *http.Request) {
	var m Message
	var su SU
	if err := su.getNewUser(r.Body); err != nil {
		m.sendErrorMessage(err.Error(), w, 400)
		return
	}
	c, err := mongoDb.GetSession()
	if err != nil {
		m.sendErrorMessage(err.Error(), w, 400)
		return
	}
	hashP, err := HashPassword(su.NewUser.Password)
	if err != nil {
		m.sendErrorMessage(err.Error(), w, 400)
	}
	u, err := c.AddUser(mongoDb.User{
		Email:    su.NewUser.Email,
		Password: hashP,
	})
	if err != nil {
		m.sendErrorMessage(err.Error(), w, 400)
		return
	}
	w.WriteHeader(200)
	b, _ := json.Marshal(map[string]interface{}{
		"user": map[string]interface{}{
			"email": u.Email,
			"token": u.Token,
		},
	})
	io.WriteString(w, string(b))
}
