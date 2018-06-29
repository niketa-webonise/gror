package services

import (
	"encoding/json"

	"github.com/gror/models"
)

// IDockerService wraps the all method of services
type ICreateDockerService interface {
	InsertData(bytevalue []byte) error
}
type IGetItemDockerService interface {
	GetItem(bytevalue []byte) (models.Root, error)
}
type IGetListDockerService interface {
	GetList(bytevalue []byte) ([]string, []string, error)
}
type IUpdateDockerService interface {
	UpdateData(bytevalue []byte) error
}

// DockerServiceImpl implements the model
type InsertDataDockerServiceImpl struct {
	CreateDockerDaoImpl models.CreateDockerDao
}
type GetItemDockerServiceImpl struct {
	GetItemDockerDaoImpl models.GetDockerItemDao
}
type GetListDockerServiceImpl struct {
	GetDockerListDaoImpl models.GetDockerListDao
}
type UpdateDockerServiceImpl struct {
	UpdateDockerDaoImpl models.UpdateDockerItemDao
}

// InsertData returns the CreateDocker's error
func (s *InsertDataDockerServiceImpl) InsertData(bytevalue []byte) error {

	var rootobject models.Root
	err := json.Unmarshal(bytevalue, &rootobject)
	if err != nil {
		return err
	}
	return s.CreateDockerDaoImpl.CreateDocker(rootobject)
}

// GetItem returns the rootobject as well as error if any
func (s *GetItemDockerServiceImpl) GetItem(bytevalue []byte) (models.Root, error) {

	var rootobject models.Root
	errUnmarshal := json.Unmarshal(bytevalue, &rootobject)
	if errUnmarshal != nil {
		return rootobject, errUnmarshal
	}
	rootobject, err := s.GetItemDockerDaoImpl.GetDockerItem(rootobject)
	return rootobject, err
}

// GetList returns the array of names and objectid
func (s *GetListDockerServiceImpl) GetList(bytevalue []byte) ([]string, []string, error) {

	var rootobject models.Root
	err := json.Unmarshal(bytevalue, &rootobject)
	if err != nil {
		return nil, nil, err
	}
	names, objid := s.GetDockerListDaoImpl.GetDockerList(rootobject)
	return names, objid, nil
}

// UpdateData returns the UpdateData's error
func (s *UpdateDockerServiceImpl) UpdateData(bytevalue []byte) error {

	var rootobject models.Root
	err := json.Unmarshal(bytevalue, &rootobject)
	if err != nil {
		return err
	}
	return s.UpdateDockerDaoImpl.UpdateDocker(rootobject)
}
