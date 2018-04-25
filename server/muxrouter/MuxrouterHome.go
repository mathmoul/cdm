package muxrouter

import (
	"net/http"
	"log"
	"errors"
	"github.com/gorilla/mux"
)

// TODO handle http response in json
// TODO handle function to create routes

const api = "/api"

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

// TODO update this function
func CreateRoute(name string, method Method, path string,
	handleFunction http.HandlerFunc, protected bool, middlewares ...http.HandlerFunc) (r Route, err error) {
	return r, err
}

// Route type
// need a []Route to handle every routes
type Route struct {
	Name        string           `json:"name"`
	Method      string           `json:"method"`
	Path        string           `json:"path"`
	HandlerFunc http.HandlerFunc `json:"func"`
	Protected   bool             `json:"protected"`
}

type Routes []Route

type MyRouter struct {
	*mux.Router
}

var R *MyRouter

func Init() {
	// os.Getenv("HOME")
	router := mux.NewRouter().StrictSlash(true)
	R = &MyRouter{
		router,
	}
}

func GetRouter() (*MyRouter) {
	if R == nil {
		Init()
	}
	return R
}

// AddRoutes
// handle and array of Route and create that routes
func (r *MyRouter) AddRoute(routes *Routes) error {
	log.Println()
	for _, route := range *routes {
		if err := route.checkroute(); err != nil {
			return err
		}
		// if route is protected or if route has middlewares
		log.Println("handle => ", api+route.Path)
		r.Handle(api+route.Path, route.HandlerFunc).Methods(route.Method).Name(route.Name)
	}
	return nil
}

// checkroute function check if route is ok and have all parameters
func (r Route) checkroute() error {
	if r.Name == "" ||
		r.Method == "" ||
		r.Path == "" ||
		r.HandlerFunc == nil {
		return errors.New("issue with this route")
	}
	return nil
}
