package controllers

import (
	"docker_orchestrator/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gror/services"
	"gopkg.in/mgo.v2/bson"
)

func CreateDockerConfig(w http.ResponseWriter, r *http.Request) {

	var rootobject model.Root
	err := json.NewDecoder(r.Body).Decode(&rootobject)
	if err != nil {
		http.Error(w, "Unprocessable Entity error", http.StatusUnprocessableEntity)
		return
	}

	rootobject.ID = bson.NewObjectId()
	marshalData, err := json.Marshal(rootobject)
	if err != nil {
		http.Error(w, "Unprocessable Entity error", http.StatusUnprocessableEntity)
		return
	}

	err = services.InsertData(marshalData)
	if err != nil {
		http.Error(w, "The request could not be completed because of a conflict", http.StatusConflict)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
	}
}

func GetDockerConfig(w http.ResponseWriter, req *http.Request) {

	var rootobject model.Root
	vars := mux.Vars(req)
	if bson.IsObjectIdHex(vars["id"]) {
		rootobject.ID = bson.ObjectIdHex(vars["id"])
		marshalData, unmarshalErr := json.Marshal(rootobject)
		if unmarshalErr != nil {
			http.Error(w, "Unprocessable Entity error", http.StatusUnprocessableEntity)
			return
		}

		rootobject, err := services.GetItem(marshalData)

		if err != nil {
			http.Error(w, "Record not found of this ID:"+vars["id"], http.StatusNotFound)
			return
		} else {
			marshalResultData, unmarshalErr := json.Marshal(rootobject)
			if unmarshalErr != nil {
				http.Error(w, "Unprocessable Entity error", http.StatusUnprocessableEntity)
				return
			}
			fmt.Fprintf(w, "%s", marshalResultData)
		}
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "Invalid Id bad request", http.StatusBadRequest)
	}
}

func UpdateDockerConfig(w http.ResponseWriter, r *http.Request) {

	var rootobject model.Root
	err := json.NewDecoder(r.Body).Decode(&rootobject)
	if err != nil {
		http.Error(w, "Unprocessable Entity error", http.StatusUnprocessableEntity)
		return
	}
	params := mux.Vars(r)

	if bson.IsObjectIdHex(params["id"]) {
		rootobject.ID = bson.ObjectIdHex(params["id"])

		marshalData, unmarshalErr := json.Marshal(rootobject)
		if unmarshalErr != nil {
			http.Error(w, "Unprocessable Entity error", http.StatusUnprocessableEntity)
			return
		}
		err = services.UpdateData(marshalData)
		if err != nil {
			http.Error(w, "Record not found of this ID:"+params["id"]+" Failed to update", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
	} else {

		http.Error(w, "Invalid Id bad request", http.StatusBadRequest)
	}
}
