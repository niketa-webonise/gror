package database

import (
	"github.com/gror/servers"
	"gopkg.in/mgo.v2"
)

type server servers.ServerDemo

type DbConnInitialiser interface {
	Init() (*mgo.Database, error)
	Collection() *mgo.Collection
}
type DbConfig struct {
	Dial   string
	DbName string
}

func (dc *DbConfig) Init() (*mgo.Database, error) {
	session, err := mgo.Dial( /*"mongodb://127.0.0.1:27017/"*/ dc.Dial)

	Db := session.DB( /*"dockerDB"*/ dc.DbName)

	return Db, err
}
