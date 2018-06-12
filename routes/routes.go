package routes

import (
	"github.com/gror/servers"
)

//RouteWrapper  wraps Server struct
type RouteWrapper struct {
	Server *servers.Server
}

//CreateRoute defines the routing with specific methods
func (r *RouteWrapper) CreateRoute() {

	//get new form
	r.Server.Router.HandleFunc("/docker/config/new", r.Server.DockerController.GetDockerConfigForm()).Methods("GET")
	//get all items list
	r.Server.Router.HandleFunc("/docker/config", r.Server.DockerController.GetDockerConfigList()).Methods("GET")
	//create new item (ajax)
	r.Server.Router.HandleFunc("/docker/config", r.Server.DockerController.CreateDockerConfig()).Methods("POST")
	//get update form for object with id {id}
	r.Server.Router.HandleFunc("/docker/config/{id}", r.Server.DockerController.GetDockerConfig()).Methods("GET")
	//update the data (ajax)
	r.Server.Router.HandleFunc("/docker/config/{id}", r.Server.DockerController.UpdateDockerConfig()).Methods("PUT")
}
