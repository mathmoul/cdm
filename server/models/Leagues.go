// Model for league
package models

import (
	"gopkg.in/mgo.v2"
	"io"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"cdm/server/database/mongoDb"
)

// League Model
type League struct {
	Collection *mgo.Collection
	// TODO Update model with requests
}

func NewLeague() *League {
	return &League{
		Collection: mongoDb.GetDb().C("league"),
	}
}

func (l *League) ParseBody(body io.Reader) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, l); err != nil {
		return err
	}
	return nil
}

func (l *League) GetAll() (L []League, err error) {
	err = l.Collection.Find(nil).All(&L)
	return
}

func (l *League) GetOneById(id string) (error) {
	err := l.Collection.FindId(id).One(l)
	return err
}

func (l *League) NewLeague() (*mgo.BulkResult, error) {
	b := l.Collection.Bulk()
	b.Insert(l)
	return b.Run()
}

func (l *League) Update(id string) (*mgo.BulkResult, error) {
	b := l.Collection.Bulk()
	b.Update(bson.M{"_id": id}, l)
	return b.Run()
}

func (l *League) Delete(id string) (*mgo.BulkResult, error) {
	b := l.Collection.Bulk()
	b.Remove(bson.M{"_id": id})
	return b.Run()
}
