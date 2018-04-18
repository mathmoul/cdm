package models

import (
	"cdm/server/database"
	"cdm/server/muxrouter"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

/*
UserCredentials for auth
*/
type UserCredentials struct {
	Credentials User `json:"credentials"`
}

type IUserCredentials interface {
	GetCredentials(r io.Reader) error
}

func (c *UserCredentials) GetCredentials(r io.Reader) error {
	body, err := ioutil.ReadAll(r)
	i := muxrouter.JSON{}
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, &i); err != nil {
		return err
	}
	if l, ok := i["user"].(map[string]interface{}); ok {
		c.Credentials.Email = l["email"].(string)
		c.Credentials.PasswordHash = l["password"].(string)
	} else if l, ok := i["credentials"].(map[string]interface{}); ok {
		c.Credentials.Email = l["email"].(string)
		c.Credentials.PasswordHash = l["password"].(string)
	}
	// c.Credentials.Email = i["user"].(map[string]interface{})["email"].(string)
	// c.Credentials.PasswordHash = i["user"].(map[string]interface{})["password"].(string)
	return nil
}

/*
User Model
*/
type User struct {
	Email             string    `json:"email" bson:"email"`                         //required
	PasswordHash      string    `json:"passwordHash" bson:"passwordHash"`           //required
	Confirmed         bool      `json:"confirmed" bson:"confirmed"`                 // required
	ConfirmationToken string    `json:"confirmationToken" bson:"confirmationToken"` //defaults to ""
	CreatedAt         time.Time `json:"createdAt" bson:"createdAt"`                 //timestamp
	UpdatedAt         time.Time `json:"updatedAt" bson:"updatedAt"`                 //timestamp
}

/*
UserReturnDatas type
*/
type UserReturnDatas struct {
	Email     string `json:"email"`
	Confirmed bool   `json:"confirmed"`
	Token     string `json:"token"`
}

/*
IUser interface
*/
type IUser interface {
	IsValidPassword( /*password */ )
	SetPassword( /* password */ )
	SetConfirmationToken()
	GenerateConfirmationURL()
	GenerateJWT() string
	ToAuthJSON()

	Login()
}

/*
IsValidPassword function
*/
func (u *User) IsValidPassword(password string) bool {
	/* first hashed password , second password to test */
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(u.PasswordHash))
	return err == nil
}

/*
SetPassword function
*/
func (u *User) SetPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), 10)
	if err != nil {
		log.Fatal(err)
	}
	u.PasswordHash = string(bytes)
}

/*
SetTimestamp function
*/
func (u *User) SetTimestamp() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

/*
SetConfirmationToken function
*/
func (u *User) SetConfirmationToken() {
	u.ConfirmationToken = u.GenerateJWT()
}

/*
GenerateConfirmationURL function
*/
func (u *User) GenerateConfirmationURL() string {
	return `http://localhost:3000/confirmation/` + u.ConfirmationToken
	//TODO remove static url
}

/*
GenerateJWT function
*/
func (u *User) GenerateJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     u.Email,
		"confirmed": u.Confirmed,
	})
	tokenS, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenS
}

/*
ToAuthJSON function
*/
func (u *User) ToAuthJSON() UserReturnDatas {
	return UserReturnDatas{
		Email:     u.Email,
		Confirmed: u.Confirmed,
		Token:     u.GenerateJWT(),
	}
}

//unique Validator

/*
Login function
Probably go to login Model
*/
func (u *User) Login() error {
	var foundUser User
	session, err := database.GetSession()
	if err != nil {
		return err
	}
	c := session.Copy().DB("cdm").C("user")
	err = c.Find(bson.M{"email": u.Email}).One(&foundUser)
	log.Println(foundUser)
	if err != nil {
		return err
	}
	if !u.IsValidPassword(foundUser.PasswordHash) {
		return errors.New("Invalid Credentials")
	}
	return nil
}

/*
InsertNewUser function
*/
func (u *User) InsertNewUser() error {
	u.SetPassword()
	u.SetTimestamp()
	u.SetConfirmationToken()
	u.Confirmed = false
	session, err := database.GetSession()
	if err != nil {
		return err
	}
	b := session.Copy().DB("cdm").C("user").Bulk()
	b.Insert(u)
	_, err = b.Run()
	if mgo.IsDup(err) {
		return errors.New("Email deja utilise")
	}
	return nil
}
