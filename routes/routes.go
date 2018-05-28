package routes

import (
	"github.com/gorilla/mux"
	"github.com/gror/controllers"
)

func CreateRoute(r *mux.Router) {

	r.HandleFunc("/docker/config", controllers.CreateDockerConfig).Methods("POST")
	r.HandleFunc("/docker/config/{id}", controllers.GetDockerConfig).Methods("GET")
	r.HandleFunc("/docker/config/{id}", controllers.UpdateDockerConfig).Methods("PUT")

}
