package database

import (
	"gopkg.in/mgo.v2"
)

//DBWrapper wraps mongoDB Database
type DBWrapper struct {
	DB *mgo.Database
}

//Init method Initialise the database connection
func (s *DBWrapper) Init() error {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/")

	s.DB = session.DB("dockerDB")

	return err
}

//Collection method return the Database collection
func (s *DBWrapper) Collection() *mgo.Collection {
	return s.DB.C("dockers")
}
