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

	// type server servers.ServerDemo

	dbconfig := &database.DbConfig{
		Dial:   "mongodb://127.0.0.1:27017/",
		DbName: "dockerDB",
	}
	db, err := dbconfig.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	sr := &servers.ServerDemo{
		Db:     db,
		Router: mux.NewRouter(),
		DockerController: &controllers.DockerControllerImpl{
			DockerService: &services.DockerServiceImpl{
				DockerDaoImpl: &models.DockerDaoImpl{
					DB: db,
				},
			},
		},
	}
	r := &routes.RouteWrapper{
		Server: sr,
	}
	r.CreateRoute()
	err = http.ListenAndServe(":8080", r.Server.Router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
