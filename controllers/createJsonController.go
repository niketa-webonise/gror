package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/docker_orchestrator/model"
	"github.com/docker_orchestrator/services"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func DockerConfig(w http.ResponseWriter, r *http.Request) {

	//defer r.Body.Close()
	var rootobject model.Root
	err := json.NewDecoder(r.Body).Decode(&rootobject)
	if err != nil {
		log.Fatal(err)
		return
	}
	rootobject.ID = bson.NewObjectId()

	marshalData, _ := json.Marshal(rootobject)

	err = services.UnmarshalJsInsert(marshalData)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", marshalData)
}

func UpdateJsonObject(w http.ResponseWriter, r *http.Request) {
	// defer r.Body.Close()
	var rootobject model.Root
	err := json.NewDecoder(r.Body).Decode(&rootobject)
	if err != nil {
		log.Fatal(err)
		return
	}
	params := mux.Vars(r)
	rootobject.ID = bson.ObjectIdHex(params["id"])
	fmt.Println(rootobject.ID)
	marshalData, _ := json.Marshal(rootobject)

	err = services.UnmarshalJsUpdate(marshalData)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", marshalData)
}
