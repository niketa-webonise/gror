package services

import (
	"encoding/json"

	"github.com/gror/models"
)

type InsertDataInerface interface {
	InsertData(bytevalue []byte) error
}

type GetItemInterface interface {
	GetItem(bytevalue []byte) (models.Root, error)
}

type UpdateDataInterface interface {
	UpdateData(bytevalue []byte) error
}

type GetListInterface interface {
	GetList() ([]string, []string, error)
}

type InsertDataImpl struct {
	InsertDockerDaoImpl models.CreateDockerInterface
}

type GetItemImpl struct {
	GetDockerDaoImpl models.GetDockerItemInterface
}

type GetListImpl struct {
	GetListDockerDaoImpl models.GetDockerListInterface
}

type UpdateDataImpl struct {
	UpdateDockerDaoImpl models.UpdateDockerInterface
}

/*InsertData method unmarshal the rootobject and calls the interface method CreateDocker
that  insert the rootobject in database*/
func (s *InsertDataImpl) InsertData(bytevalue []byte) error {
	var rootobject models.Root
	err := json.Unmarshal(bytevalue, &rootobject)
	if err != nil {
		return err
	}
	return s.InsertDockerDaoImpl.CreateDocker(rootobject)
}

/*GetItem method unmarshal the rootobject and calls the interface method GetDockerItem
that  get result by its ID*/
func (s *GetItemImpl) GetItem(bytevalue []byte) (models.Root, error) {
	var rootobject models.Root
	err := json.Unmarshal(bytevalue, &rootobject)
	if err != nil {
		return rootobject, err
	}
	rootobject, err = s.GetDockerDaoImpl.GetDockerItem(rootobject)
	return rootobject, err
}

/*GetList method return names and ids from the database*/
func (s *GetListImpl) GetList() ([]string, []string, error) {
	names, ids := s.GetListDockerDaoImpl.GetDockerList()
	return names, ids, nil
}

/*UpdateData method unmarshal the rootobject and calls the interface method UpdateDocker
that  perform updates in database*/
func (s *UpdateDataImpl) UpdateData(bytevalue []byte) error {

	var rootobject models.Root
	err := json.Unmarshal(bytevalue, &rootobject)
	if err != nil {
		return err
	}
	return s.UpdateDockerDaoImpl.UpdateDocker(rootobject)
}
