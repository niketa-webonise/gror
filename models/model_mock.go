package models

import (
	"errors"
	//"encoding/json"
)

//MockedDockerDaoImpl empty struct
type MockedDockerDaoImpl struct {
}

//GetDockerList method
func (m *MockedDockerDaoImpl) GetDockerList() ([]string, []string) {
	names := []string{"Random"}
	ids := []string{"AB345BAF"}

	return names, ids
}

func (s *MockedDockerDaoImpl) CreateDocker(rootobject Root) error {
	return errors.New("Test_create")
}

func (s *MockedDockerDaoImpl) GetDockerItem(rootobject Root) (Root, error) {
	err := errors.New("Test get item")
	return rootobject, err
}

func (s *MockedDockerDaoImpl) UpdateDocker(rootobject Root) error {
	return errors.New("update error")
}
