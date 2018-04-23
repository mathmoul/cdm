package authentication

import (
	"net/http"
	"cdm/server/models"
)

var u models.Users

func init() {
	u.Collection = u.GetModelCollection()
}

/*
Authentication == login function
 */
func Authentication(w http.ResponseWriter, r *http.Request) {
	c, err := u.ParseBody(r.Body)
	if err != nil {
		// return err
	}
	fu, err := u.GetOneByEmail(c.Email)
	if err != nil {
		//return err
	}
	if err := u.Login(fu); err != nil {
		// return err
	}
	w.WriteHeader(200)
}
