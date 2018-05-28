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
	r := mux.NewRouter()
	routes.CreateRoute(r)
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
		return
	}
}
