package database

import (
	"github.com/gror/servers"
	"gopkg.in/mgo.v2"
)

// server is a struct type variable of DockerServer
type server servers.DockerServer

// DbConnInitialiser wrap the method which is use to intialize the database
type DbConnInitialiser interface {
	Init() (*mgo.Database, error)
}

// DbConfig defines the variable which is needed to initialze the database
type DbConfig struct {
	Dial   string
	DbName string
}

// Init intialize the mongoDB database and returns the database
func (dc *DbConfig) Init() (*mgo.Database, error) {
	session, err := mgo.Dial(dc.Dial)

	Db := session.DB(dc.DbName)

	return Db, err
}
