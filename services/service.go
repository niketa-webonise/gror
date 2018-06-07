package services

import (
	"encoding/json"

	"github.com/gror/models"
)

type IDockerService interface {
	InsertData(bytevalue []byte) error
	GetItem(bytevalue []byte) (models.Root, error)
	GetList(bytevalue []byte) ([]string, []string)
	UpdateData(bytevalue []byte) error
}

type DockerServiceImpl struct {
	DockerDaoImpl models.DockerDao
}

func (s *DockerServiceImpl) InsertData(bytevalue []byte) error {

	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	return s.DockerDaoImpl.CreateDocker(rootobject)
}

func (s *DockerServiceImpl) GetItem(bytevalue []byte) (models.Root, error) {

	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	rootobject, err := s.DockerDaoImpl.GetDockerItem(rootobject)
	return rootobject, err
}
func (s *DockerServiceImpl) GetList(bytevalue []byte) ([]string, []string) {

	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	names, objid := s.DockerDaoImpl.GetDockerList(rootobject)
	return names, objid
}
func (s *DockerServiceImpl) UpdateData(bytevalue []byte) error {

	var rootobject models.Root
	json.Unmarshal(bytevalue, &rootobject)
	return s.DockerDaoImpl.UpdateDocker(rootobject)
}
