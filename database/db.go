package database

import (
	"gopkg.in/mgo.v2"
)

type DBWrapper struct {
	DB *mgo.Database
}

func (s *DBWrapper) Init() error {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/")

	s.DB = session.DB("dockerDB")

	return err
}

func (s *DBWrapper) Collection() *mgo.Collection {
	return s.DB.C("dockers")
}
