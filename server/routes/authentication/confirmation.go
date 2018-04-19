package authentication

import (
	"net/http"
	"io"
	"cdm/server/muxrouter"
	"encoding/json"
	"io/ioutil"
	"errors"
	"cdm/server/models"
)

func getToken(r io.Reader) (string, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	t := muxrouter.JSON{}
	if err := json.Unmarshal(body, &t); err != nil {
		return "", err
	}
	if _, ok := t["token"].(string); !ok {
		return "", errors.New("Issue with token")
	}
	return t["token"].(string), nil
}

func Confirmation(w http.ResponseWriter, r *http.Request) {
	var c models.User
	token, err := getToken(r.Body)
	ww := muxrouter.Mhrw{ResponseWriter: w}
	if err != nil {
		ww.Error(fasterErrors(err))
		return
	}
	c.ConfirmationToken = token
	U, err := c.ConfirmConnection()
	if err != nil {
		ww.Error(fasterErrors(err))
		return
	}
	ww.Success(muxrouter.JSON{"user": U.ToAuthJSON()})
}
