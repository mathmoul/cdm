package main

import (
	"cdm/server/muxrouter"
	"cdm/server/routes/authentication"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	_ "github.com/joho/godotenv/autoload"
	"cdm/server/routes/leagues"
	_ "cdm/server/database/mongoDb"
	"cdm/server/routes/teams"
)

func main() {
	authentication.Routes("/auth")
	//books.Books("/books")
	leagues.Routes("/leagues")
	teams.Routes("/teams")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, muxrouter.GetRouter().Router))
}
