package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gror/model"
	"github.com/gror/services"
	"gopkg.in/mgo.v2/bson"
)

func CreateDocker(w http.ResponseWriter, r *http.Request) {

	//defer r.Body.Close()
	var rootobject model.Root
	err := json.NewDecoder(r.Body).Decode(&rootobject)
	if err != nil {
		log.Fatal(err)
		return
	}
	rootobject.ID = bson.NewObjectId()

	marshalData, _ := json.Marshal(rootobject)

	err = services.UnmarshalInsertData(marshalData)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", marshalData)
}

func GetDocker(w http.ResponseWriter, req *http.Request) {

	var rootobject model.Root
	vars := mux.Vars(req)
	rootobject.ID = bson.ObjectIdHex(vars["id"])

	marshalData, _ := json.Marshal(rootobject)

	rootobject, err := services.UnmarshalGetItem(marshalData)
	if err != nil {
		log.Fatal(err)
		return
	}
	marshalResultData, _ := json.Marshal(rootobject)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", marshalResultData)
}

func UpdateDocker(w http.ResponseWriter, r *http.Request) {
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

	err = services.UnmarshalUpdateData(marshalData)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", marshalData)
}
