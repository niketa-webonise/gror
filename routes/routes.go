package routes

import (
	"github.com/gorilla/mux"
	"github.com/gror/servers"
)

// RouteWrapper defines the server
type RouteWrapper struct {
	Server *servers.DockerServer
}

// CreateRoute defines the all routes of docker
func (s *RouteWrapper) CreateRoute() *mux.Router {

	s.Server.Router.HandleFunc("/", s.Server.DockerFormController.DockerForm())
	s.Server.Router.HandleFunc("/docker/config/new", s.Server.CreateDockerController.CreateDockerConfig()).Methods("POST")
	s.Server.Router.HandleFunc("/docker", s.Server.DockerListFormController.DockerListForm())
	s.Server.Router.HandleFunc("/docker/config/list", s.Server.GetDockerConfigListController.GetDockerConfigList()).Methods("GET")
	s.Server.Router.HandleFunc("/docker/config/{id}", s.Server.GetDockerConfigController.GetDockerConfig()).Methods("GET")
	s.Server.Router.HandleFunc("/docker/config/{id}", s.Server.UpdateDockerConfigController.UpdateDockerConfig()).Methods("PUT")
	return s.Server.Router
}
