package database

import (
	"cdm/server/database/mongoDb"
	"log"
)

type Model interface {
}

type DB interface{}

var Db DB

func init () {}

func GetDb () DB  {
	if Db == nil {
		s,err := mongoDb.GetSession()
		if err == nil {
			log.Fatal(err)
		}
		Db = s
	}
	return Db

}
