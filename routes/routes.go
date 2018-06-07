package routes

import (
	"github.com/gror/servers"
)

type RouteWrapper struct {
	Server *servers.ServerDemo
}

func (s *RouteWrapper) CreateRoute() {
	s.Server.Router.HandleFunc("/", s.Server.DockerController.DockerForm())
	s.Server.Router.HandleFunc("/docker/config/new", s.Server.DockerController.CreateDockerConfig()).Methods("POST")
	s.Server.Router.HandleFunc("/docker", s.Server.DockerController.DockerListForm())
	s.Server.Router.HandleFunc("/docker/config/list", s.Server.DockerController.GetDockerConfigList()).Methods("GET")
	s.Server.Router.HandleFunc("/docker/config/{id}", s.Server.DockerController.GetDockerConfig()).Methods("GET")
	s.Server.Router.HandleFunc("/docker/config/{id}", s.Server.DockerController.UpdateDockerConfig()).Methods("PUT")

}
