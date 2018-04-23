package mongoDb

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MySession struct {
	*mgo.Session
}

/*
User ...
*/
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserDatas struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Email        string
	Password     string
	Confirmation bool
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func CreateUserCollection(s *mgo.Session) {
	session := s.Copy()
	session.DB("cdm").C("user").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}

/*
GetSession ...
*/
func GetSession() (*MySession, error) {
	s, err := mgo.Dial(os.Getenv("MONGODB_URL"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	CreateUserCollection(s)
	M := &MySession{
		s,
	}
	return M, nil
}

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

// TODO Confirmation TOken

func (s *MySession) AddUser(user User) (User, error) {
	session := s.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()
	c := session.DB("cdm").C("user")

	bulk := c.Bulk()
	tk := JwtToken(user.Email)
	bulk.Insert(&UserDatas{
		Email:        user.Email,
		Password:     user.Password,
		Token:        tk,
		Confirmation: false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	r, err := bulk.Run()
	log.Println(r, err)
	if mgo.IsDup(err) {
		return User{}, errors.New("Email deja utilise")
	}
	return User{
		Email: user.Email,
		Token: tk,
	}, nil
}

func (s *MySession) TestCredentials(u User, pwd string, f func(string, string) bool) bool {
	session := s.Copy()
	defer session.Close()
	c := session.DB("cdm").C("user")
	result := UserDatas{}
	err := c.Find(bson.M{"email": u.Email}).One(&result)
	if err != nil {
		return false
	}
	return f(pwd, result.Password)
}

func (s *MySession) FindOneAndUpdate(t string) (User, error) {
	session := s.Copy()
	c := session.DB("cdm").C("user")
	var u UserDatas
	if err := c.Find(bson.M{
		"token": t,
	}).One(&u); err != nil {
		return User{}, err
	}
	u.Confirmation = true
	u.Token = ""
	u.UpdatedAt = time.Now()
	err := c.Update(bson.M{"email": u.Email}, u)
	if err != nil {
		return User{}, err
	}
	return User{
		Email: u.Email,
		Token: u.Token,
	}, nil
}

func GetDatabaseCopy() *mgo.Database {
	session, err := GetSession()
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	return session.Copy().DB("cdm")
}