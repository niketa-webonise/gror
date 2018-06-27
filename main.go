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

	// dbconfig intialize the mongoDB dial and database name
	dbconfig := &database.DbConfig{
		Dial:   "mongodb://127.0.0.1:27017/",
		DbName: "dockerDB",
	}

	// db intialize the database
	db, err := dbconfig.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	// sr intialize the DockerServer
	sr := &servers.DockerServer{
		Db:     db,
		Router: mux.NewRouter(),
		CreateDockerController: &controllers.CreateDockerControllerImpl{
			CreateDockerService: &services.InsertDataDockerServiceImpl{
				CreateDockerDaoImpl: &models.DockerDaoImpl{
					DB: db,
				},
			},
		},
		GetDockerConfigController: &controllers.GetItemDockerControllerImpl{
			GetItemDockerService: &services.GetItemDockerServiceImpl{
				GetItemDockerDaoImpl: &models.DockerDaoImpl{
					DB: db,
				},
			},
		},
		DockerFormController: &controllers.DockerListFormImpl{},
		GetDockerConfigListController: &controllers.GetListDockerControllerImpl{
			GetListDockerService: &services.GetListDockerServiceImpl{
				GetDockerListDaoImpl: &models.DockerDaoImpl{
					DB: db,
				},
			},
		},
		DockerListFormController: &controllers.DockerListFormImpl{},
		UpdateDockerConfigController: &controllers.UpdateDockerControllerImpl{
			UpdateDockerService: &services.UpdateDockerServiceImpl{
				UpdateDockerDaoImpl: &models.DockerDaoImpl{
					DB: db,
				},
			},
		},
	}
	//r assigns the server to the RouteWrapper
	r := &routes.RouteWrapper{
		Server: sr,
	}

	r.CreateRoute()
	r.Server.Router.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	err = http.ListenAndServe(":9191", r.Server.Router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
