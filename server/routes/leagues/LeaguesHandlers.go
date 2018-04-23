package leagues

import (
	"net/http"
	"github.com/gorilla/mux"
	"cdm/server/models"
)

var l models.League

func init() {
	l.Collection = l.GetModelCollection()
}

func GetAllLeagues(w http.ResponseWriter, r *http.Request) {
	// r.Context get user
	// check only league where the user is
	_, err := l.GetAll()
	if err != nil {
		// return err
	}
	// return to json L
}

// maybe middleware to get ID_S ?

func GetOneById(w http.ResponseWriter, r *http.Request) {

	// r.Context get User
	// check if user can watch this league
	vars := mux.Vars(r)
	id := vars["id"]
	err := l.GetOneById(id)
	if err != nil {
		// return err
	}
	// return to json l
}

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	// Find league id from vars
	// Context Get user
	// check if user can modify this league
	id := mux.Vars(r)["id"]
	if err := l.ParseBody(r.Body); err != nil {
		// return err
	}
	if b, err := l.Update(id); b.Modified == 1 || err != nil {
		// return err
	}
	// return success
}

func CreateLeague(w http.ResponseWriter, r *http.Request) {
	// Context get user
	if err := l.ParseBody(r.Body); err != nil {
		// return err
	}
	b, err := l.NewLeague()
	if b.Modified == 0 || err != nil {
		// return err
	}
	// return l || success

}

func DeleteLeague(w http.ResponseWriter, r *http.Request) {
	//Context get user
	id := mux.Vars(r)["id"]
	b, err := l.Delete(id)
	if b.Modified == 0 || err != nil {
		// return err
	}
	// return success
}
