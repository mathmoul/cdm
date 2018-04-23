package database

import "cmd/database/mongoDb"

type Model interface {
	Connection() Connection
}

type Connection interface{}

func init() {
	c Connection := mongoDB.GetSession()
}
