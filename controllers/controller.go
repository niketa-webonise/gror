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

var tmpl = template.Must(template.ParseGlob("view/*.html"))

type ProjectNames struct {
	Names []string
	ObjId []string
}
type DockerConfigInterface interface {
	UpdateDockerConfig() http.HandlerFunc
	CreateDockerConfig() http.HandlerFunc
	GetDockerConfig() http.HandlerFunc
	DockerForm() http.HandlerFunc
	GetDockerConfigList() http.HandlerFunc
	DockerListForm() http.HandlerFunc
}

type DockerControllerImpl struct {
	DockerService services.IDockerService
}

func (s *DockerControllerImpl) DockerListForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "DockerList.html", nil)
	}
}
func (s *DockerControllerImpl) GetDockerConfigList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rootobject models.Root
		marshalData, unmarshalErr := json.Marshal(rootobject)
		if unmarshalErr != nil {
			http.Error(w, "Unprocessable Entity error", http.StatusUnprocessableEntity)
			return
		}
		names, objid := s.DockerService.GetList(marshalData)

		p := ProjectNames{}
		p.Names = names
		p.ObjId = objid
		tmpl.ExecuteTemplate(w, "DockerList.html", p)

	}
}
func (s *DockerControllerImpl) DockerForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "DockerForm.html", nil)
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
			if err != nil {
				http.Error(w, "Record not found of this ID:"+vars["id"], http.StatusNotFound)
				return
			} else {
				tmpl.ExecuteTemplate(w, "DockerData.html", rootobject)
			}
			w.Header().Set("Content-Type", "application/json")
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
			w.Header().Set("Content-Type", "application/json")
		} else {

			http.Error(w, "Invalid Id bad request", http.StatusBadRequest)
		}
	}
}
