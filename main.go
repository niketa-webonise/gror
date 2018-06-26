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
		CreateDockerController: &controllers.CreateDockerControllerImpl{
			CreateDockerService: &services.InsertDataImpl{
				InsertDockerDaoImpl: &models.DockerDaoImpl{
					DB: dbwrapper.DB,
				},
			},
		},
		UpdateDockerController: &controllers.UpdateDockerControllerImpl{
			UpdateDockerService: &services.UpdateDataImpl{
				UpdateDockerDaoImpl: &models.DockerDaoImpl{
					DB: dbwrapper.DB,
				},
			},
		},
		GetDockerController: &controllers.GetDockerItemControllerImpl{
			GetDockerService: &services.GetItemImpl{
				GetDockerDaoImpl: &models.DockerDaoImpl{
					DB: dbwrapper.DB,
				},
			},
		},
		GetDockerFormController: &controllers.GetDockerConfigFormImpl{},
		GetDockerListController: &controllers.GetDockerListImpl{
			GetDockerListService: &services.GetListImpl{
				GetListDockerDaoImpl: &models.DockerDaoImpl{
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

	http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("view"))))
}
