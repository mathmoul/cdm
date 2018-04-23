package muxrouter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"cdm/server/database/mongoDb"
)

type C struct {
	Credentials mongoDb.User `json:"credentials"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func (c *C) getCredentials(r io.Reader) error {
// 	body, err := ioutil.ReadAll(r)
// 	if err != nil {
// 		return err
// 	}
// 	if err = json.Unmarshal(body, c); err != nil {
// 		return err
// 	}
// 	return nil
// }

func JwtToken(e string) string {
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

func (m *Message) sendSuccessMessage(w http.ResponseWriter, u mongoDb.User) {
	w.WriteHeader(200)
	b, _ := json.Marshal(map[string]interface{}{
		"user": map[string]interface{}{
			"email": u.Email,
			"token": JwtToken(u.Email),
		},
	})
	io.WriteString(w, string(b))
}

func (m *Message) sendErrorMessage(s string, w http.ResponseWriter, status int) {
	if s == "" {
		s = "Invalid credentials"
	}
	m = &Message{
		MessageError{
			MyError{
				Global: s,
			},
		},
		MessageSuccess{},
	}
	w.WriteHeader(status)
	b, _ := json.Marshal(m.MessageError)
	io.WriteString(w, string(b))
}

// func Auth(w http.ResponseWriter, r *http.Request) {
// 	/* credentials */
// 	var u = &C{}
// 	var m = new(Message)
// 	if err := u.getCredentials(r.Body); err != nil {
// 		m.sendErrorMessage(err.Error(), w, 400)
// 		return
// 	}
// 	c, err := database.GetSession()
// 	if err != nil {
// 		log.Println(err)
// 		m.sendErrorMessage(err.Error(), w, 400)
// 		return
// 	}
// 	if !c.TestCredentials(u.Credentials, u.Credentials.Password, CheckPasswordHash) {
// 		m.sendErrorMessage("Invalid credentials", w, 400)
// 		return
// 	}
// 	m.sendSuccessMessage(w, u.Credentials)
// }

type CToken struct {
	Token string `json:"token"`
}

// func Confirmation(w http.ResponseWriter, r *http.Request) {
// 	var t CToken
// 	var m Message
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		m.sendErrorMessage(err.Error(), w, 400)
// 		return
// 	}
// 	if e := json.Unmarshal(body, &t); e != nil {
// 		m.sendErrorMessage(e.Error(), w, 400)
// 		return
// 	}
// 	// User Find One and Update with Token
// 	// confirmation token to ''
// 	// confirmation to true
// 	c, err := database.GetSession()
// 	if err != nil {
// 		m.sendErrorMessage(err.Error(), w, 400)
// 	}
// 	u, err := c.FindOneAndUpdate(t.Token)
// 	if err != nil {
// 		m.sendErrorMessage(err.Error(), w, 400)
// 		return
// 	}
// 	m.sendSuccessMessage(w, u)
// }
