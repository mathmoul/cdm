package mongoDb

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

func createUserCollection(s *mgo.Session) {
	session := s.Copy()
	session.DB("cdm").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}

var Session *mgo.Session

func initDB() {
	if Session == nil {
		session, err := mgo.Dial(os.Getenv("MONGODB_URL"))
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		if err := session.Ping(); err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		Session = session
	}
}

/*
init function
create required dbs
 */
func init() {
	initDB()
	createUserCollection(Session)
}

func GetDb() (db *mgo.Database) {
	return Session.Copy().DB(os.Getenv("DB"))
}
