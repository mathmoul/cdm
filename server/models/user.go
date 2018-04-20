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

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"fmt"
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
	ID                bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email             string        `json:"email" bson:"email"`                         //required
	PasswordHash      string        `json:"passwordHash" bson:"passwordHash"`           //required
	Confirmed         bool          `json:"confirmed" bson:"confirmed"`                 // required
	ConfirmationToken string        `json:"confirmationToken" bson:"confirmationToken"` //defaults to ""
	CreatedAt         time.Time     `json:"createdAt" bson:"createdAt"`                 //timestamp
	UpdatedAt         time.Time     `json:"updatedAt" bson:"updatedAt"`                 //timestamp
	IUser
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
	IsValidPassword(password string) bool
	SetPassword()
	SetTimestamp()
	SetConfirmationToken()
	GenerateConfirmationURL() string
	GeneratePasswordLink() string
	GenerateJWT() string
	ValidateToken() error
	GenerateResetPasswordToken() string
	ToAuthJSON() UserReturnDatas

	FindWithId(string) error

	Update() error

	Login() error
	InsertNewUser() error
	ConfirmConnection() (User, error)
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

func (u *User) GeneratePasswordLink() string {
	return `http://localhost:3000/reset_password/` + u.GenerateResetPasswordToken()
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

//TODO change name
type Z struct {
	ID bson.ObjectId `json:"_id"`
	jwt.StandardClaims
}

func (u *User) ValidateToken() error {
	token, err := jwt.ParseWithClaims(u.ConfirmationToken, &Z{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(*Z); ok && token.Valid {
		// TODO add compare claims.id with database
		return nil
	}
	return errors.New("Wrong Token")

}

func (u *User) GenerateResetPasswordToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Z{
		ID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * 365 * 24 * time.Hour).Unix(),
		},
	})
	tokenS, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatal(err)
	}

	// TODO save it into db for comparaison
	return tokenS
}

/*
ToAuthJSON function
*/
func (u *User) ToAuthJSON() UserReturnDatas {
	fmt.Println(u.Confirmed)
	return UserReturnDatas{
		Email:     u.Email,
		Confirmed: u.Confirmed,
		Token:     u.GenerateJWT(),
	}
}

//unique Validator

func findId(dt string) bson.ObjectId {
	token, err := jwt.ParseWithClaims(dt, &Z{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return ""
	}
	if claims, ok := token.Claims.(*Z); ok && token.Valid {
		// TODO add compare claims.id with database
		return claims.ID
	}
	return ""
}

func (u *User) FindWithId(token string) error {
	session, err := database.GetSession()
	if err != nil {
		return err
	}
	id := findId(token)
	c := session.Copy().DB("cdm").C("user")
	c.Find(bson.M{"_id": id}).One(u)
	return nil
}

func (u *User) Update() error {
	session, err := database.GetSession()
	if err != nil {
		return err
	}
	c := session.Copy().DB("cdm").C("user")
	fmt.Println(u)
	if err := c.Update(bson.M{"_id": u.ID}, u); err != nil {
		return err
	}
	return nil
}

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
	if err != nil {
		return err
	}
	if !u.IsValidPassword(foundUser.PasswordHash) {
		return errors.New("Invalid Credentials")
	}
	*u = foundUser
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

/*
Confirm Connection function
*/
func (u *User) ConfirmConnection() (User, error) {
	var nu User
	session, err := database.GetSession()
	if err != nil {
		return User{}, err
	}
	collection := session.Copy().DB("cdm").C("user")
	if err := collection.Find(bson.M{"confirmationToken": u.ConfirmationToken}).One(&nu); err != nil {
		return User{}, err
	}
	nu.ConfirmationToken = ""
	nu.Confirmed = true
	nu.UpdatedAt = time.Now()
	if err := collection.Update(bson.M{"confirmationToken": u.ConfirmationToken}, nu); err != nil {
		return User{}, err
	}
	return nu, nil
}
