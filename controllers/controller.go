package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gror/models"
	"github.com/gror/services"
	"gopkg.in/mgo.v2/bson"
)

// ProjectNames contains fields that is to be used to return a response from GetDockerCongifList
type ProjectNames struct {
	Names []string
	ObjId []string
}

// DockerConfigInterface wraps the all methods of controller
type UpdateDockerConfigInterface interface {
	UpdateDockerConfig() http.HandlerFunc
}
type CreateDockerConfigInterface interface {
	CreateDockerConfig() http.HandlerFunc
}
type GetDockerConfigInterface interface {
	GetDockerConfig() http.HandlerFunc
}
type DockerFormInterface interface {
	DockerForm() http.HandlerFunc
}
type GetDockerConfigListInterface interface {
	GetDockerConfigList() http.HandlerFunc
}

// DockerControllerImpl  implements the all services
type CreateDockerControllerImpl struct {
	CreateDockerService services.ICreateDockerService
}
type GetItemDockerControllerImpl struct {
	GetItemDockerService services.IGetItemDockerService
}
type GetListDockerControllerImpl struct {
	GetListDockerService services.IGetListDockerService
}
type UpdateDockerControllerImpl struct {
	UpdateDockerService services.IUpdateDockerService
}
type DockerListFormImpl struct {
}

//Absolute path
var home = os.Getenv("HOME")

// GetDockerConfigList display the DockerList page data from database and returns ProjectNames object as a response
func (s *GetListDockerControllerImpl) GetDockerConfigList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rootobject models.Root
		marshalData, unmarshalErr := json.Marshal(rootobject)
		if unmarshalErr != nil {
			http.Error(w, "Unprocessable Entity error", http.StatusUnprocessableEntity)
			return
		}
		names, objid, BadRequestErr := s.GetListDockerService.GetList(marshalData)
		if BadRequestErr != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		p := ProjectNames{}
		p.Names = names
		p.ObjId = objid
		tmpl, err := template.ParseFiles(home + "/view/DockerList.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, p)

	}
}

// DockerForm  display DockerForm page
func (s *DockerListFormImpl) DockerForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(home + "/view/DockerForm.html")
		if err != nil {
			fmt.Println(errors.New("unable to execute the template"))
		}
		tmpl.Execute(w, nil)
	}
}

// CreateDockerConfig use in POST request
func (s *CreateDockerControllerImpl) CreateDockerConfig() http.HandlerFunc {
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

		err = s.CreateDockerService.InsertData(marshalData)
		if err != nil {
			http.Error(w, "The request could not be completed because of a conflict", http.StatusConflict)
			return
		}
	}
}

// GetDockerConfig use in GET request for single record
func (s *GetItemDockerControllerImpl) GetDockerConfig() http.HandlerFunc {
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

			rootobject, err := s.GetItemDockerService.GetItem(marshalData)
			if err != nil {
				http.Error(w, "Record not found of this ID:"+vars["id"], http.StatusNotFound)
				return
			} else {
				tmpl, err := template.ParseFiles(home + "/view/DockerData.html")
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				tmpl.Execute(w, rootobject)
			}
		} else {
			http.Error(w, "Invalid Id bad request", http.StatusBadRequest)
		}
	}
}

// UpdateDockerConfig use in PUT request
func (s *UpdateDockerControllerImpl) UpdateDockerConfig() http.HandlerFunc {
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
			err = s.UpdateDockerService.UpdateData(marshalData)
			if err != nil {
				http.Error(w, "Record not found of this ID:"+params["id"]+" Failed to update", http.StatusNotFound)
				return
			}
		} else {

			http.Error(w, "Invalid Id bad request", http.StatusBadRequest)
		}
	}
}
