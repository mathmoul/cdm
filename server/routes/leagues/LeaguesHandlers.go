package leagues

import (
	"net/http"
	"github.com/gorilla/mux"
	"cdm/server/muxrouter"
	"cdm/server/models"
)

func GetAllLeagues(w http.ResponseWriter, r *http.Request) {
	l := models.NewLeague()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	// r.Context get user
	// check only league where the user is
	L, err := l.GetAll()
	if err != nil {
		j.Error401(err)
		return
	}
	j.Success(muxrouter.JSON{"data": L})
}

// maybe middleware to get ID_S ?

func GetOneById(w http.ResponseWriter, r *http.Request) {
	l := models.NewLeague()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	// r.Context get User
	// check if user can watch this league
	vars := mux.Vars(r)
	id := vars["id"]
	err := l.GetOneById(id)
	if err != nil {
		j.Error401(err)
		return
	}
	j.Success(muxrouter.JSON{"data": l})
	// return to json l
}

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	l := models.NewLeague()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	// Find league id from vars
	// Context Get user
	// check if user can modify this league
	id := mux.Vars(r)["id"]
	if err := l.ParseBody(r.Body); err != nil {
		j.Error401(err)
		return
	}
	if b, err := l.Update(id); b.Modified == 1 || err != nil {
		j.Error401(err)
		return
	}
	j.Success(nil)
}

func CreateLeague(w http.ResponseWriter, r *http.Request) {
	l := models.NewLeague()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	// Context get user
	if err := l.ParseBody(r.Body); err != nil {
		j.Error401(err)
		return
	}
	b, err := l.NewLeague()
	if b.Modified == 0 || err != nil {
		j.Error401(err)
		return
	}
	j.Success(muxrouter.JSON{"data": l})
}

func DeleteLeague(w http.ResponseWriter, r *http.Request) {
	l := models.NewLeague()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	//Context get user
	id := mux.Vars(r)["id"]
	b, err := l.Delete(id)
	if b.Modified == 0 || err != nil {
		j.Error401(err)
		return
	}
	j.Success(nil)
	// return success
}
