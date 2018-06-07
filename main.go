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
	r.Server.Router.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	err = http.ListenAndServe(":9090", r.Server.Router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
