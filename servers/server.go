package servers

import (
	"github.com/gorilla/mux"
	"github.com/gror/controllers"
	mgo "gopkg.in/mgo.v2"
)

type ServerDemo struct {
	Db               *mgo.Database
	Router           *mux.Router
	DockerController controllers.DockerConfigInterface
	// DockerService    services.IDockerService
	// DockerModel      models.DockerDao

	// DockerService    services.IDockerService
	// DockerModel      models.DockerDao
}
