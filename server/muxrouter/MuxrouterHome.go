package muxrouter

import "net/http"

// TODO handle http response in json
// TODO handle function to create routes

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

// TODO update this function
func CreateRoute(name string, method Method, path string, handleFunction http.HandlerFunc, protected bool, middlewares ...http.HandlerFunc) (error) {
	return nil
}
