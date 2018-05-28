package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	boom "github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
	"github.com/gror/model"
	"github.com/gror/services"
	"gopkg.in/mgo.v2/bson"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err1 := json.Marshal(payload)
	if err1 != nil {
		log.Fatal(err1)
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
		respondWithJson(w, http.StatusCreated, marshalData)
		respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
	}
}
func GetDockerConfig(w http.ResponseWriter, req *http.Request) {

	var rootobject model.Root
	vars := mux.Vars(req)
	if bson.IsObjectIdHex(vars["id"]) {
		rootobject.ID = bson.ObjectIdHex(vars["id"])
		marshalData, err1 := json.Marshal(rootobject)
		if err1 != nil {
			log.Fatal(err1)
			return
		}

		rootobject, err := services.GetItem(marshalData)

		if err != nil {
			boom.NotFound(w, "Record not found")
			return
		} else {
			marshalResultData, err2 := json.Marshal(rootobject)
			if err2 != nil {
				log.Fatal(err2)
				return
			}
			fmt.Fprintf(w, "%s", marshalResultData)
		}
		w.Header().Set("Content-Type", "application/json")
	} else {
		boom.BadRequest(w, "Invalid Id bad request")
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

		marshalData, err1 := json.Marshal(rootobject)
		if err1 != nil {
			log.Fatal(err1)
			return
		}
		err = services.UpdateData(marshalData)
		if err != nil {
			boom.NotFound(w, "Record not found failed to update")
			return
		} else {
			respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

		}
		w.Header().Set("Content-Type", "application/json")
	} else {
		boom.BadRequest(w, "Invalid id bad request")
	}
}
