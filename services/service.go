package services

import (
	"encoding/json"

	"github.com/gror/model"
)

func InsertData(bytevalue []byte) error {

	var rootobject model.Root
	json.Unmarshal(bytevalue, &rootobject)
	return model.CreateDocker(rootobject)
}

func GetItem(bytevalue []byte) (model.Root, error) {

	var rootobject model.Root
	json.Unmarshal(bytevalue, &rootobject)
	rootobject, err := model.GetDockerItem(rootobject)
	return rootobject, err
}

func UpdateData(bytevalue []byte) error {

	var rootobject model.Root
	json.Unmarshal(bytevalue, &rootobject)
	return model.UpdateDocker(rootobject)
}
