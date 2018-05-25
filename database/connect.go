package database

import (
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func Init() error {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/")

	db = session.DB("demoDB")

	return err
}

func Collection() *mgo.Collection {
	return db.C("samples")
}
