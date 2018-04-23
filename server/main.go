package main

import (
	"cdm/server/muxrouter"
	"cdm/server/routes/authentication"
	"cdm/server/routes/books"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	authentication.Authentication("/auth")
	books.Books("/books")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, muxrouter.GetRouter().Router))
}
