package muxrouter

import (
	"cdm/server/database"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type C struct {
	Credentials database.User `json="credentials"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (c *C) getCredentials(r io.Reader) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, c); err != nil {
		return err
	}
	return nil
}

func jwtToken(e string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": e,
	})
	fmt.Println(os.Getenv("JWT_SECRET"))
	tokenS, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenS
}

func (m *Message) sendSuccessMessage(w http.ResponseWriter, u database.User) {
	w.WriteHeader(200)
	b, _ := json.Marshal(map[string]interface{}{
		"user": map[string]interface{}{
			"email": u.Email,
			"token": jwtToken(u.Email),
		},
	})
	io.WriteString(w, string(b))
}

func (m *Message) sendErrorMessage(s string, w http.ResponseWriter, status int) {
	m = &Message{
		MessageError{
			MyError{
				Global: "Invalid credentials",
			},
		},
		MessageSuccess{},
	}
	w.WriteHeader(status)
	b, _ := json.Marshal(m.MessageError)
	io.WriteString(w, string(b))
}

func Auth(w http.ResponseWriter, r *http.Request) {
	/* credentials */
	var u = &C{}
	var m = new(Message)
	if err := u.getCredentials(r.Body); err != nil {
		m.sendErrorMessage(err.Error(), w, 400)
		return
	}
	c, err := database.GetSession()
	if err != nil {
		m.sendErrorMessage(err.Error(), w, 400)
	}
	if !c.TestCredentials(u.Credentials, u.Credentials.Password, CheckPasswordHash) {
		m.sendErrorMessage("Invalid credentials", w, 400)
		return
	}
	m.sendSuccessMessage(w, u.Credentials)
}
