package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gror/models"
	"github.com/gror/services"
	"gopkg.in/mgo.v2/bson"
)

func CreateDockerConfig(w http.ResponseWriter, r *http.Request) {

	var rootobject model.Root
	err := json.NewDecoder(r.Body).Decode(&rootobject)
	if err != nil {
		boom.BadData(w, "Unprocessable Entity error")
		return
	}

	rootobject.ID = bson.NewObjectId()

	marshalData, err := json.Marshal(rootobject)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = services.InsertData(marshalData)
	if err != nil {
		boom.Conflict(w, "The request could not be completed because of a conflict")
		return
	} else {
		w.Header().Set("Content-Type", "application/json")

	}
}
