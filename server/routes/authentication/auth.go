package authentication

import (
	"cdm/server/models"
	"cdm/server/muxrouter"
	"net/http"
)

/*
Auth function
Handler for /api/auth post request
*/
func Auth(w http.ResponseWriter, r *http.Request) {
	var c models.UserCredentials
	ww := muxrouter.Mhrw{ResponseWriter: w}
	if err := c.GetCredentials(r.Body); err != nil {
		ww.Error(muxrouter.JSON{"errors": muxrouter.JSON{"global": err.Error()}})
		return
	}
	if err := c.Credentials.Login(); err != nil {
		ww.Error(map[string]interface{}{"errors": map[string]interface{}{
			"global": err.Error(),
		}})
		return
	}
	ww.Success(muxrouter.JSON{"user": c.Credentials.ToAuthJSON()})
}
