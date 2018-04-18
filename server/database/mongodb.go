package database

import (
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
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
}

type UserDatas struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Email     string
	Password  string
	Timestamp time.Time
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
	M := &MySession{
		s,
	}
	return M, nil
}

func (s *MySession) AddUser(user User) error {
	session := s.Copy()
	defer session.Close()
	c := session.DB("cdm").C("user")
	err := c.Insert(&UserDatas{
		Email:     user.Email,
		Password:  user.Password,
		Timestamp: time.Now(),
	})
	if err != nil {
		if mgo.IsDup(err) {
			log.Println(err)
			return err
		}
	}
	return nil
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
