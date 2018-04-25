package teams

import (
	"net/http"
	"cdm/server/models"
	"cdm/server/muxrouter"
)

func GetAllTeams(w http.ResponseWriter, r *http.Request) {
	t := models.NewTeam()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	T, err := t.GetAll()
	if err != nil {
		j.Error401(err)
		return
	}
	j.Success(muxrouter.JSON{"teams": T})
}
