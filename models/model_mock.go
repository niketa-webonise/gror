package models

import (
	"errors"
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

//CreateDocker ...
func (m *MockedDockerDaoImpl) CreateDocker(rootobject Root) error {
	return errors.New("Test_create")
}

//GetDockerItem ...
func (m *MockedDockerDaoImpl) GetDockerItem(rootobject Root) (Root, error) {
	err := errors.New("Test get item")
	return rootobject, err
}

//UpdateDocker ...
func (m *MockedDockerDaoImpl) UpdateDocker(rootobject Root) error {
	return errors.New("update error")
}
