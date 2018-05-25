package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/niketa/docker_orchestrator/model"
	"github.com/niketa/docker_orchestrator/services"
	"gopkg.in/mgo.v2/bson"
)

func CreateJsonObject(w http.ResponseWriter, r *http.Request) {

	//defer r.Body.Close()
	var jsonobject model.JsonObject
	err := json.NewDecoder(r.Body).Decode(&jsonobject)
	if err != nil {
		return
	}
	jsonobject.ID = bson.NewObjectId()
	/*	err = model.InsertJsonObject(jsonobject)
		if err != nil {
			log.Fatal(err)
			return
		}*/
	marshaljb, _ := json.Marshal(jsonobject)

	err = services.Unmarshaljs(marshaljb)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", marshaljb)
}
