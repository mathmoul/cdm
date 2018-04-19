package authentication

import (
	"net/http"
	"io"
	"io/ioutil"
	"cdm/server/muxrouter"
	"encoding/json"
	"errors"
	"cdm/server/database"
	"gopkg.in/mgo.v2/bson"
	"cdm/server/models"
	"cdm/server/mailer"
	"log"
	"fmt"
)

func getEmail(reader io.Reader) (string, error) {
	var i muxrouter.JSON
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &i); err != nil {
		return "", err
	}
	email, ok := i["email"].(string)
	if !ok {
		return "", errors.New("Issue with email")
	}
	return email, nil
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var u models.User
	var mailer mailer.Mailer
	mail, err := getEmail(r.Body)
	ww := muxrouter.Mhrw{ResponseWriter: w}
	if err != nil {
		ww.Error(fasterErrors(err))
		return
	}
	s, err := database.GetSession()
	if err != nil {
		ww.Error(fasterErrors(err))
		return
	}
	c := s.Copy().DB("cdm").C("user")
	log.Println(mail)
	if err := c.Find(bson.M{"email": mail}).One(&u); err != nil {
		ww.Error(fasterErrors(errors.New("Email inconnu")))
		return
	}
	if err := mailer.SendResetPasswordEmail(u); err != nil {
		ww.Error(fasterErrors(err))
	}
	fmt.Println(u.GeneratePasswordLink())
	ww.Success(nil)
}
