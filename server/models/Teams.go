package models

import (
	"gopkg.in/mgo.v2"
	"cdm/server/database/mongoDb"
)

type Team struct {
	Collection *mgo.Collection
}

func NewTeam() *Team {
	return &Team{
		Collection: mongoDb.GetDb().C("team"),
	}
}

func (t *Team) GetAll() (T []Team, err error) {
	err = t.Collection.Find(nil).All(&T)
	if T == nil {
		T = []Team{}
	}
	return
}
