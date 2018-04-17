package muxrouter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string           `json:"name"`
	Method      string           `json:"method"`
	Path        string           `json:"path"`
	HandlerFunc http.HandlerFunc `json:"func"`
	Protected   bool             `json:"protected"`
}

type Routes []Route

type MatchaRouter struct {
	*mux.Router
}

var R *MatchaRouter

type MyError struct {
	Global string `json:"global"`
}

type Message struct {
	MyError `json:"errors"`
}

func Init() {
	_ = os.Getenv("HOME")
	router := mux.NewRouter().StrictSlash(true)
	R = &MatchaRouter{
		router,
	}
	R.AddRoute(Routes{
		Route{
			Name:   "login",
			Method: "POST",
			Path:   "/auth",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(2 * time.Second)
				m := &Message{
					MyError{
						Global: "Invalid credentials",
					},
				}
				w.WriteHeader(400)
				b, _ := json.Marshal(m)
				io.WriteString(w, string(b))
			},
			Protected: false,
		},
	})
}

func GetRouter() *MatchaRouter {
	if R == nil {
		Init()
	}
	return R
}

func (r *MatchaRouter) AddRoute(routes Routes) error {
	fmt.Println(routes)
	for _, route := range routes {
		if err := route.checkroute(); err != nil {
			return err
		}
		r.Handle("/api"+route.Path, route.HandlerFunc).Methods(route.Method).Name(route.Name)
	}
	return nil
}

func (r Route) checkroute() error {
	if r.Name == "" ||
		r.Method == "" ||
		r.Path == "" ||
		r.HandlerFunc == nil {
		return errors.New("issue with this route")
	}
	return nil
}
