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

// tmpl is used to store the all parsed html pages
var tmpl = template.Must(template.ParseGlob("view/*.html"))

// ProjectNames contains fields that is to be used to return a response from GetDockerCongifList
type ProjectNames struct {
	Names []string
	ObjId []string
}

// DockerConfigInterface wraps the all methods of controller
type DockerConfigInterface interface {
	UpdateDockerConfig() http.HandlerFunc
	CreateDockerConfig() http.HandlerFunc
	GetDockerConfig() http.HandlerFunc
	DockerForm() http.HandlerFunc
	GetDockerConfigList() http.HandlerFunc
	DockerListForm() http.HandlerFunc
}

// DockerControllerImpl  implements the all services
type DockerControllerImpl struct {
	DockerService services.IDockerService
}

// DockerListForm  display DockerList page
func (s *DockerControllerImpl) DockerListForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "DockerList.html", nil)
	}
}

// GetDockerConfigList display the DockerList page data from database and returns ProjectNames object as a response
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

// DockerForm  display DockerForm page
func (s *DockerControllerImpl) DockerForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "DockerForm.html", nil)
	}
}

// CreateDockerConfig use in POST request
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

// GetDockerConfig use in GET request for single record
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

// UpdateDockerConfig use in PUT request
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
