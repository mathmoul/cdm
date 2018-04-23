package models

import (
	"gopkg.in/mgo.v2"
	"cdm/server/database/mongoDb"
)

type Model struct{}

type IModel interface {
	GetModelCollection() (*mgo.Collection, error)

	GetAll() ([]interface{}, error)
	GetDb() *mgo.Database
}

func init() {}


func (m *Model) GetDb() *mgo.Database {
	return mongoDb.GetDatabaseCopy()
}
