package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niketa/docker_orchestrator/controllers"
)

func CreateRoute() {

	r := mux.NewRouter()

	r.HandleFunc("/createJson", controllers.CreateJsonObject).Methods("POST")
	err := http.ListenAndServe(":3500", r)
	if err != nil {
		//log.Fatal(err)
	}

}
