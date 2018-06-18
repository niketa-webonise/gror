package servers

import (
	"github.com/gorilla/mux"
	"github.com/gror/controllers"
	"github.com/gror/database"
)

//Server is a wrapper that wraps mongoDB database,mux router and DockerController interface
type Server struct {
	DB               *database.DBWrapper
	Router           *mux.Router
	DockerController controllers.DockerConfigInterface
}
