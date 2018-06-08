package servers

import (
	"github.com/gorilla/mux"
	"github.com/gror/controllers"
	mgo "gopkg.in/mgo.v2"
)

// DockerServer defines the server with database and router
type DockerServer struct {
	Db               *mgo.Database
	Router           *mux.Router
	DockerController controllers.DockerConfigInterface
}
