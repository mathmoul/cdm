package main

import (
	"cdm/server/muxrouter"
	"net/http"

	"os"

	"github.com/gorilla/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, muxrouter.GetRouter().Router))
}
