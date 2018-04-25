package middlewares

import (
	"cdm/server/muxrouter"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

/*
Authenticate function
*/
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("here")
		ww := muxrouter.Mhrw{ResponseWriter: w}
		header := r.Header.Get("authorization")
		if header == "" {
			ww.Error401(fasterErrors(errors.New("No authorisation token")))
			return
		}
		token := strings.Split(header, " ")[1]
		if token == "" {
			ww.Error401(fasterErrors(errors.New("No authorization token")))
		}
		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Invalid Token")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			ww.Error401(fasterErrors(err))
			return
		}

		// TODO add user from token.email
		next.ServeHTTP(w, r)
	})
}
