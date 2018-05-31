package routes

import (
	"github.com/gror/servers"
)

type RouteWrapper struct {
	Server *servers.Server
}

func (s *RouteWrapper) CreateRoute() {
	s.Server.Router.HandleFunc("/docker/config", s.Server.DockerController.CreateDockerConfig()).Methods("POST")
	s.Server.Router.HandleFunc("/docker/config/{id}", s.Server.DockerController.GetDockerConfig()).Methods("GET")
	s.Server.Router.HandleFunc("/docker/config/{id}", s.Server.DockerController.UpdateDockerConfig()).Methods("PUT")
}
