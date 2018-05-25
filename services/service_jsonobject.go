package services

import (
	"encoding/json"

	"github.com/docker_orchestrator/model"
)

func UnmarshalJsInsert(bytevalue []byte) error {

	var rootobject model.Root
	json.Unmarshal(bytevalue, &rootobject)
	return model.InsertJsonObject(rootobject)
}
func UnmarshalJsUpdate(bytevalue []byte) error {

	var rootobject model.Root
	json.Unmarshal(bytevalue, &rootobject)
	return model.UpdateObject(rootobject)
}
