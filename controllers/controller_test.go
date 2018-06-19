package controllers

import (
	"testing"

	"github.com/gorilla/mux"

	"github.com/gror/servers"
	"github.com/gror/services"
)

func TestCreateDockerConfig(t *testing.T) {

	s := &servers.Server{
		DB:                  dbwrapper,
		Router:              mux.NewRouter(),
		CreateDockerService: &services.InsertDataImpl{},

		UpdateDockerService: &services.UpdateDataImpl{},

		GetDockerService:     &services.GetItemImpl{},
		GetDockerListService: &services.GetListImpl{},
	}
}
