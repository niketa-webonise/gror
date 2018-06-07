package controllers

import (
	"encoding/json"
	"html/template"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/gror/models"
	"github.com/gror/services"
	"gopkg.in/mgo.v2/bson"
)

type ConfigData struct {
	Names []string
	Id    []string
}
type DockerConfigInterface interface {
	UpdateDockerConfig() http.HandlerFunc
	CreateDockerConfig() http.HandlerFunc
	GetDockerConfig() http.HandlerFunc
	GetDockerConfigForm() http.HandlerFunc
	GetDockerConfigList() http.HandlerFunc
}

type DockerControllerImpl struct {
	DockerService services.IDockerService
}

func (s *DockerControllerImpl) GetDockerConfigForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("./views/dockerconfig.gtpl")
		t.Execute(w, nil)
	}
}

func (s *DockerControllerImpl) GetDockerConfigList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		names, ids := s.DockerService.GetList()

		configData := &ConfigData{Names: names, Id: ids}

		t, _ := template.ParseFiles("./views/dockerlist.gtpl")

		t.ExecuteTemplate(w, "dockerlist.gtpl", configData)

	}
}

func (s *DockerControllerImpl) CreateDockerConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var rootobject models.Root
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

		err = s.DockerService.InsertData(marshalData)
		if err != nil {
			http.Error(w, "The request could not be completed because of a conflict", http.StatusConflict)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
		}
	}
}

func (s *DockerControllerImpl) GetDockerConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var rootobject models.Root
		vars := mux.Vars(req)
		if bson.IsObjectIdHex(vars["id"]) {
			rootobject.ID = bson.ObjectIdHex(vars["id"])
			marshalData, unmarshalErr := json.Marshal(rootobject)
			if unmarshalErr != nil {
				http.Error(w, "Unprocessable Entity error", http.StatusUnprocessableEntity)
				return
			}

			rootobject, err := s.DockerService.GetItem(marshalData)
			t, _ := template.ParseFiles("./views/dockerconfigDetails.gtpl")
			t.ExecuteTemplate(w, "dockerconfigDetails.gtpl", rootobject)
			if err != nil {
				http.Error(w, "Record not found of this ID:"+vars["id"], http.StatusNotFound)
				return
			} /*else {
				marshalResultData, unmarshalErr := json.Marshal(rootobject)
				if unmarshalErr != nil {
					http.Error(w, "Unprocessable Entity error", http.StatusUnprocessableEntity)
					return
				}
				fmt.Fprintf(w, "%s", marshalResultData)

			}*/

		} else {
			http.Error(w, "Invalid Id bad request", http.StatusBadRequest)
		}
	}
}
func (s *DockerControllerImpl) UpdateDockerConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rootobject models.Root
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
			err = s.DockerService.UpdateData(marshalData)
			if err != nil {
				http.Error(w, "Record not found of this ID:"+params["id"]+" Failed to update", http.StatusNotFound)
				return
			}
		} else {

			http.Error(w, "Invalid Id bad request", http.StatusBadRequest)
		}
	}
}
