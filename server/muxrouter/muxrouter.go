package muxrouter

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Mhrw struct {
	http.ResponseWriter
}

type JSON map[string]interface{}

func (m *Mhrw) Success(i JSON) {
	m.WriteHeader(200)
	b, _ := json.Marshal(i)
	io.WriteString(m, string(b))
}

func (m *Mhrw) Error(i JSON) {
	m.WriteHeader(400)
	b, _ := json.Marshal(i)
	io.WriteString(m, string(b))
}
func (m *Mhrw) Error401(i JSON) {
	m.WriteHeader(http.StatusUnauthorized)
	b, _ := json.Marshal(i)
	io.WriteString(m, string(b))
}

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
	MessageError
	MessageSuccess
}

type MessageSuccess struct {
	Success bool `json="succes"`
}

type MessageError struct {
	MyError `json:"errors"`
}

func Init() {
	_ = os.Getenv("HOME")
	router := mux.NewRouter().StrictSlash(true)
	R = &MatchaRouter{
		router,
	}
}

func GetRouter() *MatchaRouter {
	if R == nil {
		Init()
	}
	return R
}

func (r *Routes) Mount(mounter string) {
	y := *r
	for k := range y {
		y[k].Path = mounter + y[k].Path
	}
	r = &y
}

func (r *MatchaRouter) AddRoute(routes *Routes) error {
	for _, route := range *routes {
		if err := route.checkroute(); err != nil {
			return err
		}
		log.Println("/api" + route.Path)
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
