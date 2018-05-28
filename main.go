package main

import (
	"log"
	"net/http"
	// "net/http"
	//
	// "github.com/gorilla/mux"
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
	r := mux.NewRouter()
	routes.CreateRoute(r)
	log.Fatal(http.ListenAndServe(":7800", r))
}
