package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/gror/models"
)

type MockCreateDao struct {
}
type MockInsertDataServiceSuccess struct {
}
type MockInsertDataServiceFail struct {
}
type MockGetItemServiceSuccess struct {
}
type MockGetItemServiceFail struct {
}
type MockGetItemServiceFailToGetRecord struct {
}
type MockGetListServiceSuccess struct {
}
type MockGetListServiceFail struct {
}
type MockUpdateSuccess struct {
}
type MockUpdateFail struct {
}
type MockServices struct {
}

var TestCreateSuccess = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "1.System should insert a data into db and return a OK status",
	Request:             "{\"id\":\"5b2787d0a9036212de6a696a\",\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\"} }",
	Response:            "nil",
	StatusCode:          200,
	URLPath:             "/docker/config/new",
	controllerInterface: MockServices{},
},
}
var TestCreateFailInvalidData = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "2.System should not insert a data into db and return Unprocessable Entity error",
	Request:             "{\"id\":\"5b2787d0a9036212de6a696a\",\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\"",
	Response:            "Unprocessable Entity error",
	StatusCode:          422,
	URLPath:             "/docker/config/new",
	controllerInterface: MockServices{},
},
}
var TestCreateFailDb = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "3.System should not insert a data into db and return Failed to save in db conflict error",
	Request:             "{\"id\":\"5b2787d0a9036212de6a696a\",\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\"} }",
	Response:            "The request could not be completed because of a conflict",
	StatusCode:          409,
	URLPath:             "/docker/config/new",
	controllerInterface: MockServices{},
},
}
var TestGetItemSuccess = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "4.System should return the data of particular id",
	Request:             "5b1e8b36a9036256ef12b6e7",
	Response:            "nil",
	StatusCode:          200,
	URLPath:             "/docker/config/5b1e8b36a9036256ef12b6e7",
	controllerInterface: MockServices{},
},
}
var TestGetItemFail = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "5.System should return Invalid Id bad request error",
	Request:             "11212",
	Response:            "Invalid Id bad request",
	StatusCode:          400,
	URLPath:             "/docker/config/11212",
	controllerInterface: MockServices{},
},
}
var TestGetItemFailToGetRecord = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "6.System should return Record not found of this ID error",
	Request:             "5b1e8b36a9036256ef12b6e7",
	Response:            "Record not found of this ID:5b1e8b36a9036256ef12b6e7",
	StatusCode:          404,
	URLPath:             "/docker/config/5b1e8b36a9036256ef12b6e7",
	controllerInterface: MockServices{},
},
}
var TestGetListSuccess = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "7.System should return the all list of project",
	Request:             "",
	Response:            "nil",
	StatusCode:          200,
	URLPath:             "/docker/config/list",
	controllerInterface: MockServices{},
},
}
var TestUpdateSuccess = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "8.System should update the data in db",
	Request:             "{\"id\":\"5b2787d0a9036212de6a696a\",\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\"} }",
	Response:            "nil",
	StatusCode:          200,
	URLPath:             "/docker/config/5b1e8b36a9036256ef12b6e7",
	controllerInterface: MockServices{},
},
}
var TestUpdateFailInvalidData = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "11.System should not update the data in db and return Unprocessable Entity error",
	Request:             "{\"id\":\"5b2787d0a9036212de6a696a\",\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\"",
	Response:            "Unprocessable Entity error",
	StatusCode:          422,
	URLPath:             "/docker/config/5b1e8b36a9036256ef12b6e7",
	controllerInterface: MockServices{},
},
}
var TestUpdateFail = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "9.System should return bad request error",
	Request:             "{\"id\":\"5b2787d0a9036212de6a696a\",\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\"} }",
	Response:            "Invalid Id bad request",
	StatusCode:          400,
	URLPath:             "/docker/config/11212",
	controllerInterface: MockServices{},
},
}
var TestUpdateFailInvalidId = []struct {
	Title               string
	Request             string
	Response            string
	StatusCode          int
	URLPath             string
	controllerInterface MockServices
}{{
	Title:               "10.System should return Record not found of this ID error",
	Request:             "{\"id\":\"5b2787d0a9036212de6a696a\",\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\"} }",
	Response:            "Record not found of this ID:5b1e8b36a9036256ef12b6e7 Failed to update",
	StatusCode:          404,
	URLPath:             "/docker/config/5b1e8b36a9036256ef12b6e7",
	controllerInterface: MockServices{},
},
}

func (s MockInsertDataServiceSuccess) InsertData(bytevalue []byte) error {
	return nil
}
func (s MockInsertDataServiceFail) InsertData(bytevalue []byte) error {
	return errors.New("The request could not be completed because of a conflict")
}
func (s MockGetItemServiceSuccess) GetItem(bytevalue []byte) (models.Root, error) {
	var rootobject models.Root
	return rootobject, nil
}
func (s MockGetItemServiceFail) GetItem(bytevalue []byte) (models.Root, error) {
	var rootobject models.Root
	return rootobject, errors.New("Bad request")
}
func (s MockGetItemServiceFailToGetRecord) GetItem(bytevalue []byte) (models.Root, error) {
	var rootobject models.Root
	return rootobject, errors.New("Record not found of this ID:5b1e8b36a9036256ef12b6e7")
}
func (s MockGetListServiceSuccess) GetList(bytevalue []byte) ([]string, []string, error) {

	return []string{}, []string{}, nil
}
func (s MockUpdateSuccess) UpdateData(bytevalue []byte) error {
	return nil
}
func (s MockUpdateFail) UpdateData(bytevalue []byte) error {
	return errors.New("Record not found of this ID:5b1e8b36a9036256ef12b6e7")
}
func TestCreateDockerConfig(t *testing.T) {

	for _, testCase := range TestCreateSuccess {
		r := mux.NewRouter()
		server := httptest.NewServer(mux.NewRouter())

		t.Log("\n")
		t.Log("Executing test", testCase.Title)

		c := &CreateDockerControllerImpl{
			CreateDockerService: MockInsertDataServiceSuccess{},
		}
		reader := strings.NewReader(testCase.Request)

		r.HandleFunc("/docker/config/new", c.CreateDockerConfig())
		//Create new request
		request, err := http.NewRequest("POST", testCase.URLPath, reader)
		if err != nil {
			fmt.Println(errors.New("Failed to create new request"))
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		//Satisfy handlerfunc method
		r.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}
		server.Close()

	}
	for _, testCase := range TestCreateFailInvalidData {

		r := mux.NewRouter()
		server := httptest.NewServer(mux.NewRouter())

		t.Log("\n")
		t.Log("Executing test", testCase.Title)

		c := &CreateDockerControllerImpl{
			CreateDockerService: MockInsertDataServiceFail{},
		}
		reader := strings.NewReader(testCase.Request)

		r.HandleFunc("/docker/config/new", c.CreateDockerConfig())
		//Create new request
		request, err := http.NewRequest("POST", testCase.URLPath, reader)
		if err != nil {
			fmt.Println(errors.New("Failed to create new request"))
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		//Satisfy handlerfunc method
		r.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}

		//validate the error messages
		if strings.TrimSpace(string(w.Body.Bytes())) != testCase.Response {

			t.Logf("expected message to be %s but got %s", testCase.Response, string(w.Body.Bytes()))
			t.Fail()
			continue
		}
		server.Close()

	}
	for _, testCase := range TestCreateFailDb {

		r := mux.NewRouter()
		server := httptest.NewServer(mux.NewRouter())

		t.Log("\n")
		t.Log("Executing test", testCase.Title)

		c := &CreateDockerControllerImpl{
			CreateDockerService: MockInsertDataServiceFail{},
		}
		reader := strings.NewReader(testCase.Request)

		r.HandleFunc("/docker/config/new", c.CreateDockerConfig())
		//Create new request
		request, err := http.NewRequest("POST", testCase.URLPath, reader)
		if err != nil {
			fmt.Println(errors.New("Failed to create new request"))
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		//Satisfy handlerfunc method
		r.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}

		//validate the error messages
		if strings.TrimSpace(string(w.Body.Bytes())) != testCase.Response {

			t.Logf("expected message to be %s but got %s", testCase.Response, string(w.Body.Bytes()))
			t.Fail()
			continue
		}
		server.Close()

	}

}
func TestGetDockerConfig(t *testing.T) {

	for _, testCase := range TestGetItemSuccess {
		r := mux.NewRouter()
		t.Log("\n")
		t.Log("Executing test", testCase.Title)
		c := &GetItemDockerControllerImpl{
			GetItemDockerService: MockGetItemServiceSuccess{},
		}
		r.HandleFunc("/docker/config/{id}", c.GetDockerConfig())

		server := httptest.NewServer(r)
		defer server.Close()
		//Create new request
		request, err := http.NewRequest("GET", testCase.URLPath, nil)

		if err != nil {
			t.Fatal("Failed to create new request")
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		//Satisfy handlerfunc method
		r.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}
	}
	for _, testCase := range TestGetItemFail {
		r := mux.NewRouter()
		t.Log("\n")
		t.Log("Executing test", testCase.Title)
		c := &GetItemDockerControllerImpl{
			GetItemDockerService: MockGetItemServiceFail{},
		}
		r.HandleFunc("/docker/config/{id}", c.GetDockerConfig())

		server := httptest.NewServer(r)
		defer server.Close()
		//Create new request
		request, err := http.NewRequest("GET", testCase.URLPath, nil)

		if err != nil {
			t.Fatal("Failed to create new request")
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		//Satisfy handlerfunc method
		r.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}

		//validate the error messages
		if strings.TrimSpace(string(w.Body.Bytes())) != testCase.Response {

			t.Logf("expected message to be %s but got %s", testCase.Response, string(w.Body.Bytes()))
			t.Fail()
			continue
		}

	}
	for _, testCase := range TestGetItemFailToGetRecord {
		r := mux.NewRouter()
		t.Log("\n")
		t.Log("Executing test", testCase.Title)
		c := &GetItemDockerControllerImpl{
			GetItemDockerService: MockGetItemServiceFailToGetRecord{},
		}
		r.HandleFunc("/docker/config/{id}", c.GetDockerConfig())

		server := httptest.NewServer(r)
		defer server.Close()
		//Create new request
		request, err := http.NewRequest("GET", testCase.URLPath, nil)

		if err != nil {
			t.Fatal("Failed to create new request")
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		//Satisfy handlerfunc method
		r.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}

		//validate the error messages
		if strings.TrimSpace(string(w.Body.Bytes())) != testCase.Response {

			t.Logf("expected message to be %s but got %s", testCase.Response, string(w.Body.Bytes()))
			t.Fail()
			continue
		}

	}

}
func TestGetDockerConfigList(t *testing.T) {

	server := httptest.NewServer(mux.NewRouter())

	for _, testCase := range TestGetListSuccess {

		t.Log("\n")
		t.Log("Executing test", testCase.Title)

		reader := strings.NewReader(testCase.Request)

		//Create new request
		request, err := http.NewRequest("GET", testCase.URLPath, reader)
		if err != nil {
			fmt.Println(errors.New("Failed to create new request"))
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		c := &GetListDockerControllerImpl{
			GetListDockerService: MockGetListServiceSuccess{},
		}

		handler := http.HandlerFunc(c.GetDockerConfigList())

		//Satisfy handlerfunc method
		handler.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}
	}
	server.Close()
}
func TestUpdateDockerConfig(t *testing.T) {

	for _, testCase := range TestUpdateSuccess {
		r := mux.NewRouter()
		t.Log("\n")
		t.Log("Executing test", testCase.Title)
		c := &UpdateDockerControllerImpl{
			UpdateDockerService: MockUpdateSuccess{},
		}
		r.HandleFunc("/docker/config/{id}", c.UpdateDockerConfig())

		server := httptest.NewServer(r)
		defer server.Close()
		reader := strings.NewReader(testCase.Request)
		//Create new request
		request, err := http.NewRequest("PUT", testCase.URLPath, reader)

		if err != nil {
			t.Fatal("Failed to create new request")
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		//Satisfy handlerfunc method
		r.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}
	}
	for _, testCase := range TestUpdateFail {
		r := mux.NewRouter()
		t.Log("\n")
		t.Log("Executing test", testCase.Title)
		c := &UpdateDockerControllerImpl{
			UpdateDockerService: MockUpdateFail{},
		}
		r.HandleFunc("/docker/config/{id}", c.UpdateDockerConfig())

		server := httptest.NewServer(r)
		defer server.Close()
		reader := strings.NewReader(testCase.Request)

		//Create new request
		request, err := http.NewRequest("PUT", testCase.URLPath, reader)

		if err != nil {
			t.Fatal("Failed to create new request")
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		//Satisfy handlerfunc method
		r.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}

		//validate the error messages
		if strings.TrimSpace(string(w.Body.Bytes())) != testCase.Response {

			t.Logf("expected message to be %s but got %s", testCase.Response, string(w.Body.Bytes()))
			t.Fail()
			continue
		}

	}
	for _, testCase := range TestUpdateFailInvalidId {
		r := mux.NewRouter()
		t.Log("\n")
		t.Log("Executing test", testCase.Title)
		c := &UpdateDockerControllerImpl{
			UpdateDockerService: MockUpdateFail{},
		}
		r.HandleFunc("/docker/config/{id}", c.UpdateDockerConfig())

		server := httptest.NewServer(r)
		defer server.Close()
		reader := strings.NewReader(testCase.Request)

		//Create new request
		request, err := http.NewRequest("PUT", testCase.URLPath, reader)

		if err != nil {
			t.Fatal("Failed to create new request")
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		//Satisfy handlerfunc method
		r.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}

		//validate the error messages
		if strings.TrimSpace(string(w.Body.Bytes())) != testCase.Response {

			t.Logf("expected message to be %s but got %s", testCase.Response, string(w.Body.Bytes()))
			t.Fail()
			continue
		}

	}
	for _, testCase := range TestUpdateFailInvalidData {
		r := mux.NewRouter()
		t.Log("\n")
		t.Log("Executing test", testCase.Title)
		c := &UpdateDockerControllerImpl{
			UpdateDockerService: MockUpdateFail{},
		}
		r.HandleFunc("/docker/config/{id}", c.UpdateDockerConfig())

		server := httptest.NewServer(r)
		defer server.Close()
		reader := strings.NewReader(testCase.Request)

		//Create new request
		request, err := http.NewRequest("GET", testCase.URLPath, reader)

		if err != nil {
			t.Fatal("Failed to create new request")
		}

		w := httptest.NewRecorder()

		t.Log(testCase.Title)

		//Satisfy handlerfunc method
		r.ServeHTTP(w, request)

		//validate the API codes
		if w.Code != testCase.StatusCode {
			t.Logf("got code %d but expected %d", w.Code, testCase.StatusCode)
			t.Fail()
			continue
		}

		//validate the error messages
		if strings.TrimSpace(string(w.Body.Bytes())) != testCase.Response {

			t.Logf("expected message to be %s but got %s", testCase.Response, string(w.Body.Bytes()))
			t.Fail()
			continue
		}

	}

}
