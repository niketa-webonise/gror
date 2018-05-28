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
		log.Fatal(err)
		return
	}
	rootobject.ID = bson.NewObjectId()

	marshalData, err1 := json.Marshal(rootobject)
	if err1 != nil {
		log.Fatal(err1)
		return
	}
	err = services.InsertData(marshalData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Fail to insert")
		return
	} else {
		marshalData, err1 := json.Marshal(rootobject)
		if err1 != nil {
			log.Fatal(err1)
			return
		}
		respondWithJson(w, http.StatusCreated, marshalData)
		respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

	}
	w.Header().Set("Content-Type", "application/json")
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

		rootobject, err := services.UnmarshalGetItem(marshalData)

		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Record not found")
			return
		} else {
			marshalResultData, _ := json.Marshal(rootobject)
			fmt.Fprintf(w, "%s", marshalResultData)

		}

		w.Header().Set("Content-Type", "application/json")

	} else {
		respondWithError(w, http.StatusBadRequest, "Invalid id bad request")
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
		err = services.UnmarshalUpdateData(marshalData)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Record not found fail to update")
			return
		} else {
			respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

		}
		w.Header().Set("Content-Type", "application/json")
	} else {
		respondWithError(w, http.StatusBadRequest, "Invalid id bad request")
	}

}
