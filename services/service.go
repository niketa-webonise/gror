package services

import (
	"encoding/json"

	"github.com/gror/models"
)

//IDockerService interface wraps all methods of services
type IDockerService interface {
	InsertData(bytevalue []byte) error
	GetItem(bytevalue []byte) (models.Root, error)
	UpdateData(bytevalue []byte) error
	GetList() ([]string, []string)
}

//DockerServiceImpl defines DockerDaoImpl which is of type DockerDao interface
type DockerServiceImpl struct {
	DockerDaoImpl models.DockerDao
}

/*InsertData method unmarshal the rootobject and calls the interface method CreateDocker
that  insert the rootobject in database*/
func (s *DockerServiceImpl) InsertData(bytevalue []byte) error {
	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	return s.DockerDaoImpl.CreateDocker(rootobject)
}

/*GetItem method unmarshal the rootobject and calls the interface method GetDockerItem
that  get result by its ID*/
func (s *DockerServiceImpl) GetItem(bytevalue []byte) (models.Root, error) {
	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	rootobject, err := s.DockerDaoImpl.GetDockerItem(rootobject)
	return rootobject, err
}

/*GetList method return names and ids from the database*/
func (s *DockerServiceImpl) GetList() ([]string, []string) {

	names, ids := s.DockerDaoImpl.GetDockerList()
	return names, ids
}

/*UpdateData method unmarshal the rootobject and calls the interface method UpdateDocker
that  perform updates in database*/
func (s *DockerServiceImpl) UpdateData(bytevalue []byte) error {

	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	return s.DockerDaoImpl.UpdateDocker(rootobject)
}
