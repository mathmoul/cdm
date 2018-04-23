package main

import (
	"cdm/server/muxrouter"
	"cdm/server/routes/authentication"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	_ "github.com/joho/godotenv/autoload"
	"cdm/server/routes/leagues"
)

func main() {
	authentication.AuthenticationRoutes("/auth")
	//books.Books("/books")
	leagues.LeaguesRoutes("/leagues")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, muxrouter.GetRouter().Router))
}
