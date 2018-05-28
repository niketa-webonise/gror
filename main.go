package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gror/database"
	"github.com/gror/routes"
)

func main() {

	err := database.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	router := mux.NewRouter()
	routes.CreateRoute(router)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
		return
	}

}
