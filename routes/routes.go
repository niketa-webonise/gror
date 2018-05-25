package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gror/controllers"
)

func CreateRoute() {

	r := mux.NewRouter()

	r.HandleFunc("/docker/config", controllers.CreateDocker).Methods("POST")
	r.HandleFunc("/docker/config/{id}", controllers.GetDocker).Methods("GET")
	r.HandleFunc("/docker/config/{id}", controllers.UpdateDocker).Methods("PUT")
	log.Fatal(http.ListenAndServe(":7800", r))

}
