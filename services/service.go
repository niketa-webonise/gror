package services

import (
	"encoding/json"

	"github.com/gror/models"
)

// IDockerService wraps the all method of services
type IDockerService interface {
	InsertData(bytevalue []byte) error
	GetItem(bytevalue []byte) (models.Root, error)
	GetList(bytevalue []byte) ([]string, []string)
	UpdateData(bytevalue []byte) error
}

// DockerServiceImpl implements the model
type DockerServiceImpl struct {
	DockerDaoImpl models.DockerDao
}

// InsertData returns the CreateDocker's error
func (s *DockerServiceImpl) InsertData(bytevalue []byte) error {

	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	return s.DockerDaoImpl.CreateDocker(rootobject)
}

// GetItem returns the rootobject as well as error if any
func (s *DockerServiceImpl) GetItem(bytevalue []byte) (models.Root, error) {

	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	rootobject, err := s.DockerDaoImpl.GetDockerItem(rootobject)
	return rootobject, err
}

// GetList returns the array of names and objectid
func (s *DockerServiceImpl) GetList(bytevalue []byte) ([]string, []string) {

	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	names, objid := s.DockerDaoImpl.GetDockerList(rootobject)
	return names, objid
}

// UpdateData returns the UpdateData's error
func (s *DockerServiceImpl) UpdateData(bytevalue []byte) error {

	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	return s.DockerDaoImpl.UpdateDocker(rootobject)
}
