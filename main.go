package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gror/controllers"
	"github.com/gror/database"
	"github.com/gror/models"
	"github.com/gror/routes"
	"github.com/gror/servers"

	"github.com/gror/services"
)

func main() {

	dbwrapper := &database.DBWrapper{}
	dbwrapper.Init()
	s := &servers.Server{
		DB:     dbwrapper,
		Router: mux.NewRouter(),
		DockerController: &controllers.DockerControllerImpl{
			DockerService: &services.DockerServiceImpl{
				DockerDaoImpl: &models.DockerDaoImpl{
					DB: dbwrapper.DB,
				},
			},
		},
	}
	r := &routes.RouteWrapper{
		Server: s,
	}
	r.CreateRoute()
	err := http.ListenAndServe(":8080", r.Server.Router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
