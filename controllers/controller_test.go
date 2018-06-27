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

type UpdateIDSuccessImpl struct {
}

type UpdateIDFailImpl struct {
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

var testCaseUpdateIDSuccess = []struct {
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
	names := []string{"cacoon-cam", "cacooncam"}
	ids := []string{"5b2cd0a9a90362508f80f71d", "5b28b442a90362768113e47e"}
	return names, ids, nil
}

func (s UpdateIDFullSuccessImpl) UpdateData(bytevalue []byte) error {
	return nil
}

func (s UpdateIDSuccessImpl) UpdateData(bytevalue []byte) error {
	return errors.New("valid ID ,invalid JSONBody")
}

func (s UpdateIDFailImpl) UpdateData(bytevalue []byte) error {
	return errors.New("invalid ID")
}

func TestCreateDockerConfig(t *testing.T) {

	s := CreateDockerControllerImpl{
		CreateDockerService: CreateDockerSuccessImplTest{},
	}

	for _, test := range testCaseCreateSuccess {
		router := mux.NewRouter()
		ts := httptest.NewServer(router)
		defer ts.Close()

		req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.JSONBody))

		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(s.CreateDockerConfig())
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

	s = CreateDockerControllerImpl{
		CreateDockerService: CreateDockerFailImplTest{},
	}

	for _, test := range testCaseCreateFail {
		router := mux.NewRouter()
		ts := httptest.NewServer(router)
		defer ts.Close()

		req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.JSONBody))
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(s.CreateDockerConfig())
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

	s := GetDockerItemControllerImpl{
		GetDockerService: GetItemIDFailImplTest{},
	}

	for _, test := range testCaseGetItemIDFail {

		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", s.GetDockerConfig())
		ts := httptest.NewServer(router)
		defer ts.Close()

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

	s = GetDockerItemControllerImpl{
		GetDockerService: GetItemReqSuccessImplTest{},
	}

	for _, test := range testCaseGetItemFullSuccess {
		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", s.GetDockerConfig())

		ts := httptest.NewServer(router)
		defer ts.Close()

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

	s := UpdateDockerControllerImpl{
		UpdateDockerService: UpdateIDFullSuccessImpl{},
	}

	for _, test := range testCaseUpdateIDFullSuccess {

		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", s.UpdateDockerConfig())

		ts := httptest.NewServer(router)
		defer ts.Close()

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

	s = UpdateDockerControllerImpl{
		UpdateDockerService: UpdateIDSuccessImpl{},
	}

	for _, test := range testCaseUpdateIDSuccess {

		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", s.UpdateDockerConfig())

		ts := httptest.NewServer(router)
		defer ts.Close()

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

	s = UpdateDockerControllerImpl{
		UpdateDockerService: UpdateIDFailImpl{},
	}

	for _, test := range testCaseUpdateIDFail {

		router := mux.NewRouter()
		router.HandleFunc("/docker/config/{id}", s.UpdateDockerConfig())

		ts := httptest.NewServer(router)
		defer ts.Close()

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

}
func TestGetList(t *testing.T) {

	s := GetDockerListImpl{
		GetDockerListService: GetListSuccessImplTest{},
	}

	for _, test := range testCaseGetListSuccess {
		router := mux.NewRouter()
		ts := httptest.NewServer(router)
		defer ts.Close()
		req, err := http.NewRequest(test.Method, test.URL, nil)

		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(s.GetDockerConfigList())
		handler.ServeHTTP(resp, req)

		if status := resp.Code; status != test.expectStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectStatus)
		}
	}
}
