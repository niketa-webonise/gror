package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type CreateDockerImplTest struct {
}

var testCaseCreateFail = []struct {
	URL       string
	Message   string
	Name      string
	expectErr error
}{
	{
		URL:       "/docker/config",
		Message:   "successfully created",
		expectErr: nil,
	},
}

func (s CreateDockerImplTest) CreateDockerConfig() error {
	return nil

}
func TestCreateDockerConfig(t *testing.T) {

	router := mux.NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	r := &CreateDockerControllerImpl{},

}
