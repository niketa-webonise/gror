package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"os"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/gror/models"
	"github.com/gror/services"
	"gopkg.in/mgo.v2/bson"
)

//ConfigData defines Names and Id fields
type ConfigData struct {
	Names []string //The name of the project
	ID    []string //The unique Id store for each record in database
}

//UpdateDockerConfigInterface interface
type UpdateDockerConfigInterface interface {
	UpdateDockerConfig() http.HandlerFunc
}

//CreateDockerConfigInterface interface
type CreateDockerConfigInterface interface {
	CreateDockerConfig() http.HandlerFunc
}

//GetDockerConfigInterface interface
type GetDockerConfigInterface interface {
	GetDockerConfig() http.HandlerFunc
}

//GetDockerConfigFormInterface interface
type GetDockerConfigFormInterface interface {
	GetDockerConfigForm() http.HandlerFunc
}

//GetDockerConfigListInterface interface
type GetDockerConfigListInterface interface {
	GetDockerConfigList() http.HandlerFunc
}

type UpdateDockerControllerImpl struct {
	UpdateDockerService services.UpdateDataInterface
}

type CreateDockerControllerImpl struct {
	CreateDockerService services.InsertDataInerface
}

type GetDockerItemControllerImpl struct {
	GetDockerService services.GetItemInterface
}

type GetDockerListImpl struct {
	GetDockerListService services.GetListInterface
}

type GetDockerConfigFormImpl struct {
}

//Path variable
var Path = os.Getenv("GO_PATH")

//GetDockerConfigForm method execute the template "dockerconfig.gtpl".
func (s *GetDockerConfigFormImpl) GetDockerConfigForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(Path + "views/dockerconfig.gtpl")
		if err != nil {
			fmt.Println(errors.New("unable to execute the template"))
		}
		t.Execute(w, nil)
	}
}

/*GetDockerConfigList method execute the template "dockerlist.gtpl"
and in response sending the struct that contains names and ids.*/
func (s *GetDockerListImpl) GetDockerConfigList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		names, ids, err := s.GetDockerListService.GetList()
		if err != nil {
			fmt.Println(errors.New("Unable to get Ids and names"))
		}

		configData := &ConfigData{Names: names, ID: ids}

		t, err := template.ParseFiles(Path + "views/dockerlist.gtpl")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		t.ExecuteTemplate(w, "dockerlist.gtpl", configData)
	}
}

//CreateDockerConfig method get called on POST request and return response in Header
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
		} else {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, "{\"message\":\"Successfully created!\"}")
		}

	}
}

//GetDockerConfig method get called on GET request
func (s *GetDockerItemControllerImpl) GetDockerConfig() http.HandlerFunc {
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

			rootobject, err := s.GetDockerService.GetItem(marshalData)
			if err != nil {
				http.Error(w, "Record not found of this ID:"+vars["id"], http.StatusNotFound)
				return
			}

			t, err := template.ParseFiles(Path + "views/dockerconfigDetails.gtpl")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			t.ExecuteTemplate(w, "dockerconfigDetails.gtpl", rootobject)
		} else {
			http.Error(w, "Invalid Id bad request", http.StatusBadRequest)
		}
	}
}

/*UpdateDockerConfig method get called on PUT request*/
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
		fmt.Fprintln(w, "{\"message\":\"Successfully updated!\"}")
	}
}
