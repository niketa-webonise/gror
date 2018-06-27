package servers

import (
	"github.com/gorilla/mux"
	"github.com/gror/controllers"
	mgo "gopkg.in/mgo.v2"
)

// DockerServer defines the server with database and router
type DockerServer struct {
	Db                            *mgo.Database
	Router                        *mux.Router
	CreateDockerController        controllers.CreateDockerConfigInterface
	GetDockerConfigController     controllers.GetDockerConfigInterface
	DockerFormController          controllers.DockerFormInterface
	GetDockerConfigListController controllers.GetDockerConfigListInterface
	DockerListFormController      controllers.DockerListFormInterface
	UpdateDockerConfigController  controllers.UpdateDockerConfigInterface
}
