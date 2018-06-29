package controllers

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/gror/models"
)

type CreateDockerFailImplTest struct {
}

type CreateDockerSuccessImplTest struct {
}

type GetItemIDFailImplTest struct {
}

type GetItemReqSuccessImplTest struct {
}

type GetListSuccessImplTest struct {
}

type GetListFailImplTest struct {
}

type UpdateIDFullSuccessImpl struct {
}

type UpdateFailInvalidDataImpl struct {
}

type UpdateIDFailImpl struct {
}

type UpdateInvalidIDImpl struct {
}

var testCaseCreateFail = []struct {
	Method       string
	URL          string
	Message      string
	JSONBody     string
	expectStatus int
}{
	{
		Method:       "POST",
		URL:          "/docker/config",
		Message:      "Unprocessable Entity error",
		JSONBody:     "{\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\" }",
		expectStatus: 422,
	},
}
var testCaseCreateSuccess = []struct {
	Method       string
	URL          string
	Message      string
	JSONBody     string
	expectStatus int
}{
	{
		Method:       "POST",
		URL:          "/docker/config",
		Message:      "{\"message\":\"Successfully created!\"}",
		JSONBody:     "{\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\"} }",
		expectStatus: 200,
	},
}

var testCaseGetItemIDFail = []struct {
	ID           string
	Method       string
	URL          string
	Message      string
	expectStatus int
}{
	{
		ID:           "1100",
		Method:       "GET",
		URL:          "/docker/config/1100",
		Message:      "Invalid Id bad request",
		expectStatus: 400,
	},
}

var testCaseGetItemFullSuccess = []struct {
	ID           string
	Method       string
	URL          string
	expectStatus int
}{
	{
		ID:           "5b28b442a90362768113e47e",
		Method:       "GET",
		URL:          "/docker/config/5b28b442a90362768113e47e",
		expectStatus: 200,
	},
}

var testCaseUpdateIDFullSuccess = []struct {
	ID           string
	Method       string
	Message      string
	URL          string
	JSONBody     string
	expectStatus int
}{
	{
		ID:           "5b28b442a90362768113e47e",
		Method:       "PUT",
		Message:      "{\"message\":\"Successfully updated!\"}",
		URL:          "/docker/config/5b28b442a90362768113e47e",
		JSONBody:     "{\"systemInfo\": {\"grorVersion\": \"1.1.0\",\"name\": \"cocooncam\"} }",
		expectStatus: 200,
	},
}

var testCaseUpdateInvalidData = []struct {
	ID           string
	Method       string
	Message      string
	URL          string
	JSONBody     string
	expectStatus int
}{
	{
		ID:           "5b28b442a90362768113e47e",
		Method:       "PUT",
		Message:      "Unprocessable Entity error",
		URL:          "/docker/config/5b28b442a90362768113e47e",
		JSONBody:     "{\"systemInfo\": {\"grorVersion\": \"1.1.0\",\"name\": \"cocooncam\" }",
		expectStatus: 422,
	},
}

var testCaseUpdateIDFail = []struct {
	ID           string
	Message      string
	Method       string
	URL          string
	JSONBody     string
	expectStatus int
}{
	{
		ID:           "12345",
		Message:      "Unprocessable Entity error",
		Method:       "PUT",
		URL:          "/docker/config/12345",
		JSONBody:     "{\"systemInfo\": {\"grorVersion\": \"1.1.0\",\"name\": \"cocooncam\" }",
		expectStatus: 422,
	},
}

var testCaseUpdateInvalidID = []struct {
	ID           string
	Message      string
	Method       string
	URL          string
	JSONBody     string
	expectStatus int
}{
	{
		ID:           "5b28b442a90362768113e47e",
		Message:      "Record not found of this ID:5b28b442a90362768113e47e Failed to update",
		Method:       "PUT",
		URL:          "/docker/config/5b28b442a90362768113e47e",
		JSONBody:     "{\"systemInfo\": {\"grorVersion\": \"1.1.0\",\"name\": \"cocooncam\" } }",
		expectStatus: 404,
	},
}

var testCaseGetListSuccess = []struct {
	Method       string
	URL          string
	expectStatus int
}{
	{
		Method:       "GET",
		URL:          "/docker/config",
		expectStatus: 200,
	},
}

func (s CreateDockerSuccessImplTest) InsertData(bytevalue []byte) error {
	return nil
}

func (s CreateDockerFailImplTest) InsertData(bytevalue []byte) error {
	return errors.New("The request could not be completed because of a conflict")
}

func (s GetItemIDFailImplTest) GetItem(bytevalue []byte) (models.Root, error) {
	var rootobject models.Root
	return rootobject, errors.New("id not found")
}

func (s GetItemReqSuccessImplTest) GetItem(bytevalue []byte) (models.Root, error) {
	var rootobject models.Root
	return rootobject, nil
}

func (s GetListSuccessImplTest) GetList() ([]string, []string, error) {
	return []string{}, []string{}, nil
}

func (s UpdateIDFullSuccessImpl) UpdateData(bytevalue []byte) error {
	return nil
}

func (s UpdateFailInvalidDataImpl) UpdateData(bytevalue []byte) error {
	return errors.New("valid ID ,invalid JSONBody")
}

func (s UpdateIDFailImpl) UpdateData(bytevalue []byte) error {
	return errors.New("invalid ID")
}

func (s UpdateInvalidIDImpl) UpdateData(bytevalue []byte) error {
	return errors.New("Record not found of this ID")
}

func TestCreateDockerConfig(t *testing.T) {
	for _, test := range testCaseCreateSuccess {
		c := CreateDockerControllerImpl{
			CreateDockerService: CreateDockerSuccessImplTest{},
		}
		router := mux.NewRouter()
		server := httptest.NewServer(router)
		defer server.Close()

		req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.JSONBody))

		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(c.CreateDockerConfig())
		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(resp, req)
		requestBody := string(resp.Body.Bytes())

		if strings.TrimSpace(requestBody) != test.Message {
			t.Errorf("expected message to be %s but got %s", test.Message, requestBody)
		}

		if status := resp.Code; status != test.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectStatus)
		}
	}

	for _, test := range testCaseCreateFail {

		c := CreateDockerControllerImpl{
			CreateDockerService: CreateDockerFailImplTest{},
		}

		router := mux.NewRouter()
		server := httptest.NewServer(router)
		defer server.Close()

		req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.JSONBody))
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(c.CreateDockerConfig())
		handler.ServeHTTP(resp, req)

		requestBody := string(resp.Body.Bytes())
		if strings.TrimSpace(requestBody) != test.Message {
			t.Errorf("expected message to be %s but got %s", test.Message, requestBody)
		}

		if status := resp.Code; status != test.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectStatus)
		}
	}
}

func TestGetItem(t *testing.T) {

	for _, test := range testCaseGetItemIDFail {

		c := GetDockerItemControllerImpl{
			GetDockerService: GetItemIDFailImplTest{},
		}

		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", c.GetDockerConfig())
		server := httptest.NewServer(router)
		defer server.Close()

		req, err := http.NewRequest(test.Method, test.URL, nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		requestBody := string(resp.Body.Bytes())

		if strings.TrimSpace(requestBody) != test.Message {
			t.Errorf("expected message to be %s but got %s", test.Message, requestBody)
		}

		IDval, idErr := hex.DecodeString(test.ID)
		if status := resp.Code; status != test.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		} else {
			if idErr != nil {
				fmt.Printf("Invalid ID:%s", hex.EncodeToString(IDval))
			}
		}
	}

	for _, test := range testCaseGetItemFullSuccess {
		c := GetDockerItemControllerImpl{
			GetDockerService: GetItemReqSuccessImplTest{},
		}

		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", c.GetDockerConfig())

		server := httptest.NewServer(router)
		defer server.Close()

		req, err := http.NewRequest(test.Method, test.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		IDval, idErr := hex.DecodeString(test.ID)
		if status := resp.Code; status != test.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectStatus)
		} else {
			if idErr != nil {
				fmt.Printf("Invalid ID:%s", hex.EncodeToString(IDval))
			}
		}
	}
}

func TestUpdateData(t *testing.T) {

	c := UpdateDockerControllerImpl{
		UpdateDockerService: UpdateIDFullSuccessImpl{},
	}

	for _, test := range testCaseUpdateIDFullSuccess {

		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", c.UpdateDockerConfig())

		server := httptest.NewServer(router)
		defer server.Close()

		req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.JSONBody))
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		requestBody := string(resp.Body.Bytes())

		if strings.TrimSpace(requestBody) != test.Message {
			t.Errorf("expected message to be %s but got %s", test.Message, requestBody)
		}

		if status := resp.Code; status != test.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectStatus)
		}
	}

	for _, test := range testCaseUpdateInvalidData {
		c := UpdateDockerControllerImpl{
			UpdateDockerService: UpdateInvalidIDImpl{},
		}

		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", c.UpdateDockerConfig())

		server := httptest.NewServer(router)
		defer server.Close()

		req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.JSONBody))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		requestBody := string(resp.Body.Bytes())

		if strings.TrimSpace(requestBody) != test.Message {
			t.Errorf("expected message to be %s but got %s", test.Message, requestBody)
		}

		if status := resp.Code; status != test.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectStatus)
		}
	}

	for _, test := range testCaseUpdateIDFail {

		c := UpdateDockerControllerImpl{
			UpdateDockerService: UpdateIDFailImpl{},
		}

		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", c.UpdateDockerConfig())

		server := httptest.NewServer(router)
		defer server.Close()

		req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.JSONBody))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		requestBody := string(resp.Body.Bytes())

		if strings.TrimSpace(requestBody) != test.Message {
			t.Errorf("expected message to be %s but got %s", test.Message, requestBody)
		}

		if status := resp.Code; status != test.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectStatus)
		}
	}

	for _, test := range testCaseUpdateInvalidID {

		c := UpdateDockerControllerImpl{
			UpdateDockerService: UpdateInvalidIDImpl{},
		}

		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", c.UpdateDockerConfig())

		server := httptest.NewServer(router)
		defer server.Close()

		reader := strings.NewReader(test.JSONBody)

		req, err := http.NewRequest(test.Method, test.URL, reader)
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		requestBody := string(resp.Body.Bytes())

		if strings.TrimSpace(requestBody) != test.Message {
			t.Errorf("expected message to be %s but got %s", test.Message, requestBody)
		}

		if status := resp.Code; status != test.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectStatus)
		}
	}

}
func TestGetList(t *testing.T) {
	for _, test := range testCaseGetListSuccess {
		c := GetDockerListImpl{
			GetDockerListService: GetListSuccessImplTest{},
		}
		router := mux.NewRouter()
		server := httptest.NewServer(router)
		defer server.Close()
		req, err := http.NewRequest(test.Method, test.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(c.GetDockerConfigList())
		handler.ServeHTTP(resp, req)

		if status := resp.Code; status != test.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectStatus)
		}
	}
}
