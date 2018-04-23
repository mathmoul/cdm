package models

import (
	"cdm/server/database/mongoDb"
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
UsersCredentials for auth
*/
/*
json :
{
	"credentials":
		{
			"email":"mm@mm.com",
			"password":"1234"
		}
}
 */
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
Users Model
*/
type Users struct {
	ID                bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email             string        `json:"email" bson:"email"`                         //required
	PasswordHash      string        `json:"passwordHash" bson:"passwordHash"`           //required
	Confirmed         bool          `json:"confirmed" bson:"confirmed"`                 // required
	ConfirmationToken string        `json:"confirmationToken" bson:"confirmationToken"` //defaults to ""
	CreatedAt         time.Time     `json:"createdAt" bson:"createdAt"`                 //timestamp
	UpdatedAt         time.Time     `json:"updatedAt" bson:"updatedAt"`                 //timestamp
	IUsers
	Model
	Collection        *mgo.Collection
}

/*
UsersReturnDatas type
*/
type UsersReturnDatas struct {
	Email     string `json:"email"`
	Confirmed bool   `json:"confirmed"`
	Token     string `json:"token"`
}

/*
IUsers interface
*/
type IUsers interface {
	isValidPassword(password string) bool
	SetPassword()
	SetTimestamp()
	SetConfirmationToken()
	GenerateConfirmationURL() string
	GeneratePasswordLink() string
	GenerateJWT() string
	ValidateToken() error
	GenerateResetPasswordToken() string
	ToAuthJSON() UsersReturnDatas

	FindWithId(string) error

	Update() error

	Login(foundUsers Users) error
	InsertNewUsers() error
	ConfirmConnection() (Users, error)
}

// PRIVATE FUNCTIONS
/*
IsValidPassword function
*/
func (u *Users) isValidPassword(password string) bool {
	/* first hashed password , second password to test */
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(u.PasswordHash))
	return err == nil
}

/*
SetPassword function
*/
func (u *Users) SetPassword() {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), 10)
	if err != nil {
		log.Fatal(err)
	}
	u.PasswordHash = string(bytes)
}

/*
SetTimestamp function
*/
func (u *Users) SetTimestamp() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

/*
SetConfirmationToken function
*/
func (u *Users) SetConfirmationToken() {
	u.ConfirmationToken = u.GenerateJWT()
}

/*
GenerateConfirmationURL function
*/
func (u *Users) GenerateConfirmationURL() string {
	return `http://localhost:3000/confirmation/` + u.ConfirmationToken
	//TODO remove static url
}

func (u *Users) GeneratePasswordLink() string {
	return `http://localhost:3000/reset_password/` + u.GenerateResetPasswordToken()
}

/*
GenerateJWT function
*/
func (u *Users) GenerateJWT() string {
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

func (u *Users) ValidateToken() error {
	token, err := jwt.ParseWithClaims(u.ConfirmationToken, &Z{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(*Z); ok && token.Valid {
		// TODO add compare claims.id with mongoDb
		return nil
	}
	return errors.New("Wrong Token")

}

func (u *Users) GenerateResetPasswordToken() string {
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
func (u *Users) ToAuthJSON() UsersReturnDatas {
	fmt.Println(u.Confirmed)
	return UsersReturnDatas{
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
		// TODO add compare claims.id with mongoDb
		return claims.ID
	}
	return ""
}

func (u *Users) FindWithId(token string) error {
	session, err := mongoDb.GetSession()
	if err != nil {
		return err
	}
	id := findId(token)
	c := session.Copy().DB("cdm").C("user")
	c.Find(bson.M{"_id": id}).One(u)
	return nil
}

func (u *Users) Update() error {
	session, err := mongoDb.GetSession()
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
InsertNewUsers function
*/
func (u *Users) InsertNewUsers() error {
	u.SetPassword()
	u.SetTimestamp()
	u.SetConfirmationToken()
	u.Confirmed = false
	session, err := mongoDb.GetSession()
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
func (u *Users) ConfirmConnection() (Users, error) {
	var nu Users
	session, err := mongoDb.GetSession()
	if err != nil {
		return Users{}, err
	}
	collection := session.Copy().DB("cdm").C("user")
	if err := collection.Find(bson.M{"confirmationToken": u.ConfirmationToken}).One(&nu); err != nil {
		return Users{}, err
	}
	nu.ConfirmationToken = ""
	nu.Confirmed = true
	nu.UpdatedAt = time.Now()
	if err := collection.Update(bson.M{"confirmationToken": u.ConfirmationToken}, nu); err != nil {
		return Users{}, err
	}
	return nu, nil
}

//PUBLIC FUNCTIONS
/*
Login function
Probably go to login Model
*/
func (u *Users) Login(foundUsers Users) (err error) {
	if !u.isValidPassword(foundUsers.PasswordHash) {
		return errors.New("Invalid Credentials")
	}
	*u = foundUsers
	return
}

// New Functions
/*
json :
{
	"credentials":
		{
			"email":"mm@mm.com",
			"password":"1234"
		}
}
 */

func (u *Users) ParseBody(reader io.Reader) (l Credentials, err error) {
	var b []byte
	m := map[string]Credentials{
		"credentials": l,
	}
	b, err = ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return
	}
	u.PasswordHash = m["credentials"].Password
	u.Email = m["credentials"].Email
	return
}

/*
Database functions
 */

func (u *Users) GetModelCollection() *mgo.Collection {
	return u.Model.GetDb().C("users")
}

func (u *Users) GetOneByEmail(email string) (nu Users, err error) {
	err = u.Collection.Find(bson.M{"email": email}).One(&nu)
	return
}

/*
 */
