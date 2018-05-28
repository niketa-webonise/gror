package routes

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/gror/controllers"
)

func CreateRoute(r *mux.Router) {

	fmt.Println(reflect.TypeOf(r))
	r.HandleFunc("/docker/config", controllers.CreateDockerConfig).Methods("POST")
	r.HandleFunc("/docker/config/{id}", controllers.GetDockerConfig).Methods("GET")
	r.HandleFunc("/docker/config/{id}", controllers.UpdateDockerConfig).Methods("PUT")
	log.Fatal(http.ListenAndServe(":7800", r))

}
