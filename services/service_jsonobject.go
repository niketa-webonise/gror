package services

import (
	"encoding/json"

	"github.com/niketa/docker_orchestrator/model"
)

func Unmarshaljs(bytevalue []byte) error {

	var jsonObject model.JsonObject
	json.Unmarshal(bytevalue, &jsonObject)
	return model.InsertJsonObject(jsonObject)
}
