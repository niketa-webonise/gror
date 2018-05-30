package servers

import (
	"github.com/gorilla/mux"
	"github.com/gror/controllers"
	"github.com/gror/database"
)

type Server struct {
	DB               *database.DBWrapper
	Router           *mux.Router
	DockerController controllers.DockerConfigInterface
}
