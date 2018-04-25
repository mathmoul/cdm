package models

import (
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
	"cdm/server/database/mongoDb"
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
	Email             string        `json:"email" bson:"email"`               //required
	PasswordHash      string        `json:"passwordHash" bson:"passwordHash"` //required
	Confirmed         bool          `json:"confirmed" bson:"confirmed"`       // required
	Admin             bool          `json:"admin" bson:"admin"`
	ConfirmationToken string        `json:"confirmationToken" bson:"confirmationToken"` //defaults to ""
	CreatedAt         time.Time     `json:"createdAt" bson:"createdAt"`                 //timestamp
	UpdatedAt         time.Time     `json:"updatedAt" bson:"updatedAt"`                 //timestamp
	IUsers
	Collection        *mgo.Collection
}

func NewUsers() *Users {
	return &Users{
		ID:                "",
		Admin:             false,
		Email:             "",
		PasswordHash:      "",
		Collection:        mongoDb.GetDb().C("users"),
		ConfirmationToken: "",
		Confirmed:         false,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

/*
UsersReturnDatas type
*/
type UsersReturnDatas struct {
	Email     string `json:"email"`
	Confirmed bool   `json:"confirmed"`
	Admin     bool   `json:"admin"`
	Token     string `json:"token"`
}

/*
IUsers interface
*/
type IUsers interface {
	isValidPassword(password string) bool
	SetPassword()
	setTimestamp()
	setConfirmationToken()
	GenerateConfirmationURL() string
	GeneratePasswordLink() string
	generateJWT() string
	ValidateToken() error
	GenerateResetPasswordToken() string
	ToAuthJSON() UsersReturnDatas

	FindWithId(string) error

	Update() (*mgo.BulkResult, error)

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
func (u *Users) setTimestamp() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

/*
SetConfirmationToken function
*/
func (u *Users) setConfirmationToken() {
	u.ConfirmationToken = u.generateJWT()
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
func (u *Users) generateJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     u.Email,
		"confirmed": u.Confirmed,
		"admin":     u.Admin,
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
	fmt.Println(err)
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
			ExpiresAt: time.Now().Add(31 * 24 * time.Hour).Unix(), // token is valid for 31 days => one month
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
		Admin:     u.Admin,
		Token:     u.generateJWT(),
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
	return u.Collection.Find(bson.M{"_id": findId(token)}).One(u)
}

func (u *Users) Update() (*mgo.BulkResult, error) {
	var b *mgo.Bulk

	u.UpdatedAt = time.Now()
	b = u.Collection.Bulk()
	b.Update(bson.M{"_id": u.ID}, u)
	return b.Run()
}

/*
Confirm Connection function
*/
func (u *Users) ConfirmConnection() (nu Users, err error) {
	if err = u.Collection.Find(bson.M{"confirmationToken": u.ConfirmationToken}).One(&nu); err != nil {
		return
	}
	nu.ConfirmationToken = ""
	nu.Confirmed = true
	nu.UpdatedAt = time.Now()
	b := u.Collection.Bulk()
	b.Update(bson.M{"confirmationToken": u.ConfirmationToken}, nu)
	_, err = b.Run()
	return
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

func (u *Users) InsertNewUsers() (err error) {
	u.SetPassword()
	u.setTimestamp()
	u.setConfirmationToken()
	u.Confirmed = false
	if _, err = u.NewUsers(); mgo.IsDup(err) {
		return
	}
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

func (u *Users) ParseBody(name string, reader io.Reader) (l Credentials, err error) {
	var b []byte
	m := map[string]Credentials{
		name: l,
	}
	b, err = ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return
	}
	u.PasswordHash = m[name].Password
	l.Password = m[name].Password
	u.Email = m[name].Email
	l.Email = m[name].Email
	return
}

/*
Database functions
 */

func (u *Users) GetOneByEmail(email string) (nu Users, err error) {
	err = u.Collection.Find(bson.M{"email": email}).One(&nu)
	return
}

func (u *Users) NewUsers() (*mgo.BulkResult, error) {
	b := u.Collection.Bulk()
	b.Insert(u)
	return b.Run()
}

/*
 */
