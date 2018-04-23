package books

import (
	"cdm/server/muxrouter"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"cdm/server/routes/middlewares"
)

const (
	confKey int = iota
)

func Books(mounter string) {
	r := &muxrouter.Routes{
		muxrouter.Route{
			Name:        "GetAll",
			Method:      "GET",
			Path:        mounter,
			HandlerFunc: middlewares.Authenticate(GetAll),
			Protected:   false,
		},
		muxrouter.Route{
			Name:        "search",
			Method:      "GET",
			Path:        mounter + "/search",
			HandlerFunc: middlewares.Authenticate(SearchBook),
			Protected:   false,
		},
	}
	muxrouter.GetRouter().AddRoute(r)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("==================> ", "GetAll")
	log.Println(r.Context())
}

type Book struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Title  string        `json:"title" bson:"title"`
	Author string        `json:"author" bson:"author"`
	Covers []string      `json:"covers" bson:"covers"`
	Pages  uint32        `json:"pages" bson:"pages"`
}

func SearchBook(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(map[string]interface{}{
		"books": []Book{
			Book{
				ID:     "5addab7f12c184012d9dd7b7",
				Title:  "sint",
				Author: "Kristy Allen",
				Covers: []string{
					"www.culpa.com/image",
					"www.id.com/image",
					"www.magna.com/image",
				},
				Pages: 474,
			},
			{
				ID:     "5addab7fb1852406ca1ecfa0",
				Title:  "cupidatat",
				Author: "Kidd Valeria",
				Covers: []string{
					"www.sint.com/image",
					"www.ea.com/image",
				},
				Pages: 991,
			},
			{
				ID:     "5addab7f5c123b768a01e49c",
				Title:  "irure",
				Author: "Stokes Schneider",
				Covers: []string{
					"www.mollit.com/image",
					"www.aute.com/image",
					"www.fugiat.com/image",
				},
				Pages: 709,
			},
			{
				ID:     "5addab7f4ae45aa37bc70415",
				Title:  "enim",
				Author: "Katina Frye",
				Covers: []string{
					"www.enim.com/image",
					"www.ut.com/image",
					"www.duis.comoimage",
				},
				Pages: 268,
			},
		},
	})
	io.WriteString(w, string(b))
}
