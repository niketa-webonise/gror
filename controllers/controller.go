package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/gror/models"
	"github.com/gror/services"
	"gopkg.in/mgo.v2/bson"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func CreateDockerConfig(w http.ResponseWriter, r *http.Request) {

	var rootobject model.Root
	err := json.NewDecoder(r.Body).Decode(&rootobject)
	if err != nil {
		boom.BadData(w, "Unprocessable Entity Error")
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
		respondWithJson(w, http.StatusCreated, map[string]string{"result": "success"})
	}
}

func GetDockerConfig(w http.ResponseWriter, req *http.Request) {

	var rootobject model.Root
	vars := mux.Vars(req)
	if bson.IsObjectIdHex(vars["id"]) {
		rootobject.ID = bson.ObjectIdHex(vars["id"])
		marshalData, unmarshalErr := json.Marshal(rootobject)
		if unmarshalErr != nil {
			log.Fatal(unmarshalErr)
			return
		}

		rootobject, err := services.GetItem(marshalData)

		if err != nil {
			boom.NotFound(w, "Data not found with this ID "+vars["id"])
			return
		} else {
			marshalResultData, unmarshalErr := json.Marshal(rootobject)
			if unmarshalErr != nil {
				log.Fatal(unmarshalErr)
				return
			}
			fmt.Fprintf(w, "%s", marshalResultData)
		}
		w.Header().Set("Content-Type", "application/json")
	} else {
		boom.BadRequest(w, "Invalid Id bad request")
		return
	}
}

func UpdateDockerConfig(w http.ResponseWriter, r *http.Request) {

	var rootobject model.Root
	err := json.NewDecoder(r.Body).Decode(&rootobject)
	if err != nil {
		log.Fatal(err)
		return
	}
	params := mux.Vars(r)

	if bson.IsObjectIdHex(params["id"]) {
		rootobject.ID = bson.ObjectIdHex(params["id"])

		marshalData, unmarshalErr := json.Marshal(rootobject)
		if unmarshalErr != nil {
			log.Fatal(unmarshalErr)
			return
		}
		err = services.UpdateData(marshalData)
		if err != nil {
			boom.NotFound(w, "Data not found with this ID "+params["id"])
			return
		} else {
			respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
		}
		w.Header().Set("Content-Type", "application/json")
	} else {
		boom.BadRequest(w, "Invalid id bad request")
		return
	}
}
