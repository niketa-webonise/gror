package routes

import (
	"log"
	"net/http"

	"github.com/docker_orchestrator/controllers"
	"github.com/gorilla/mux"
)

func CreateRoute() {

	r := mux.NewRouter()

	r.HandleFunc("/docker/config", controllers.DockerConfig).Methods("POST")
	r.HandleFunc("/docker/{id}", controllers.UpdateJsonObject).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))

}
