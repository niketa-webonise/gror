package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gror/models"
)

type CreateSuccessImplTest struct {
}
type CreateFailImplTest struct {
}
type GetItemFailImplTest struct {
}
type GetItemSuccessImplTest struct {
}
type UpdateSuccessImplTest struct {
}
type UpdateFailImplTest struct {
}
type GetListSuccessImplTest struct {
}
type CreateFailMarshalTest struct {
}

var testCaseCreateFail = []struct {
	Url       string
	Message   string
	Name      string
	expectErr error
}{
	{
		Url:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data but fails in saving data to db",
		expectErr: errors.New("Failed to save in db"),
	},
}

var testCaseCreateSuccess = []struct {
	Url              string
	Message          string
	Name             string
	expectErr        error
	serviceInterface CreateSuccessImplTest
}{
	{
		Url:              "../sample_gror.json",
		Message:          "successfully unmarshall and saved in db",
		Name:             "valid data and successfully saved in db",
		expectErr:        nil,
		serviceInterface: CreateSuccessImplTest{},
	},
}

var testCaseGetitemFail = []struct {
	Url              string
	Message          string
	Name             string
	expectErr        error
	serviceInterface GetItemFailImplTest
}{
	{
		Url:              "../sample_gror.json",
		Message:          "successfully unmarshall",
		Name:             "valid data but fails to get the item",
		expectErr:        errors.New("Failed to get the item"),
		serviceInterface: GetItemFailImplTest{},
	},
}

var testCaseGetitemSuccess = []struct {
	Url       string
	Message   string
	Name      string
	expectErr error
}{
	{
		Url:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data and get the item",
		expectErr: nil,
	},
}

var testCaseGetListSuccess = []struct {
	Url       string
	Message   string
	Name      string
	expectErr error
}{
	{
		Url:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data and get the whole list",
		expectErr: nil,
	},
}
var testCaseUpdateFail = []struct {
	Url       string
	Message   string
	Name      string
	expectErr error
}{
	{
		Url:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data but fail to update data",
		expectErr: errors.New("Failed to update data"),
	},
}
var testCaseUpdateSuccess = []struct {
	Url       string
	Message   string
	Name      string
	expectErr error
}{
	{
		Url:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data and data updated",
		expectErr: nil,
	},
}

var testCaseCreateFailMarshal = []struct {
	Url       string
	Message   string
	Name      string
	expectErr error
}{
	{
		Url:       "../sample_gror1.json",
		Message:   "unsuccessfully unmarshall",
		Name:      "invalid data",
		expectErr: errors.New("invalid data"),
	},
}

func (s CreateSuccessImplTest) CreateDocker(rootobject models.Root) error {
	return nil

}
func (s CreateFailImplTest) CreateDocker(rootobject models.Root) error {
	return errors.New("Failed to save in db")
}

func (s CreateFailMarshalTest) CreateDocker(rootobject models.Root) error {
	return errors.New("invalid data")
}

func (s GetItemFailImplTest) GetDockerItem(rootobject models.Root) (models.Root, error) {

	return rootobject, errors.New("Failed to get the item")
}
func (s GetItemSuccessImplTest) GetDockerItem(rootobject models.Root) (models.Root, error) {

	return rootobject, nil
}
func (s CreateFailMarshalTest) GetDockerItem(rootobject models.Root) (models.Root, error) {

	return rootobject, errors.New("invalid data")
}
func (s UpdateFailImplTest) UpdateDocker(rootobject models.Root) error {

	return errors.New("Failed to update data")
}
func (s UpdateSuccessImplTest) UpdateDocker(rootobject models.Root) error {

	return nil
}

func (s CreateFailMarshalTest) UpdateDocker(rootobject models.Root) error {

	return errors.New("invalid data")
}
func (s GetListSuccessImplTest) GetDockerList() ([]string, []string) {

	return []string{}, []string{}
}

func TestInsertData(t *testing.T) {
	var rootobject models.Root

	r := &InsertDataImpl{
		InsertDockerDaoImpl: CreateFailImplTest{},
	}

	for _, gror := range testCaseCreateFail {

		raw, err := ioutil.ReadFile(gror.Url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &rootobject)

		actualErr := r.InsertData(raw)
		if gror.expectErr.Error() != actualErr.Error() {
			panic("test case failed for " + gror.Name)
		}
	}
	r = &InsertDataImpl{
		InsertDockerDaoImpl: CreateSuccessImplTest{},
	}

	for _, gror := range testCaseCreateSuccess {

		raw, err := ioutil.ReadFile(gror.Url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &rootobject)

		actualErr := r.InsertData(raw)

		if gror.expectErr != actualErr {
			//t.Errorf("error= %q, want %q", err1, gror.Message)
			panic("test case failed for " + gror.Name)
		}
	}

	r = &InsertDataImpl{
		InsertDockerDaoImpl: CreateFailMarshalTest{},
	}

	for _, gror := range testCaseCreateFailMarshal {

		raw, err := ioutil.ReadFile(gror.Url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &rootobject)

		actualErr := r.InsertData(raw)

		if gror.expectErr.Error() != actualErr.Error() {
			panic("test case failed for " + gror.Name)
		}
	}

}

func TestGetItem(t *testing.T) {
	var rootobject models.Root

	r := &GetItemImpl{
		GetDockerDaoImpl: GetItemFailImplTest{},
	}

	for _, gror := range testCaseGetitemFail {

		raw, err := ioutil.ReadFile(gror.Url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &rootobject)

		_, actualErr := r.GetItem(raw)
		if gror.expectErr.Error() != actualErr.Error() {
			panic("test case failed for " + gror.Name)
		}
	}
	r = &GetItemImpl{
		GetDockerDaoImpl: GetItemSuccessImplTest{},
	}

	for _, gror := range testCaseGetListSuccess {

		raw, err := ioutil.ReadFile(gror.Url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &rootobject)

		_, actualErr := r.GetItem(raw)

		if gror.expectErr != actualErr {
			//t.Errorf("error= %q, want %q", err1, gror.Message)
			panic("test case failed for " + gror.Name)
		}
	}

	r = &GetItemImpl{
		GetDockerDaoImpl: CreateFailMarshalTest{},
	}

	for _, gror := range testCaseCreateFailMarshal {

		raw, err := ioutil.ReadFile(gror.Url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &rootobject)

		_, actualErr := r.GetItem(raw)

		if gror.expectErr.Error() != actualErr.Error() {
			panic("test case failed for " + gror.Name)
		}
	}

}

func TestUpdateData(t *testing.T) {
	var rootobject models.Root

	r := &UpdateDataImpl{
		UpdateDockerDaoImpl: UpdateFailImplTest{},
	}

	for _, gror := range testCaseUpdateFail {

		raw, err := ioutil.ReadFile(gror.Url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &rootobject)

		actualErr := r.UpdateData(raw)
		if gror.expectErr.Error() != actualErr.Error() {
			panic("test case failed for " + gror.Name)
		}
	}
	r = &UpdateDataImpl{
		UpdateDockerDaoImpl: UpdateSuccessImplTest{},
	}

	for _, gror := range testCaseUpdateSuccess {

		raw, err := ioutil.ReadFile(gror.Url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &rootobject)

		actualErr := r.UpdateData(raw)

		if gror.expectErr != actualErr {
			panic("test case failed for " + gror.Name)
		}
	}

	r = &UpdateDataImpl{
		UpdateDockerDaoImpl: CreateFailMarshalTest{},
	}

	for _, gror := range testCaseCreateFailMarshal {

		raw, err := ioutil.ReadFile(gror.Url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &rootobject)

		actualErr := r.UpdateData(raw)

		if gror.expectErr.Error() != actualErr.Error() {
			panic("test case failed for " + gror.Name)
		}
	}
}

func TestGetList(t *testing.T) {
	var rootobject models.Root

	r := &GetListImpl{
		GetListDockerDaoImpl: GetListSuccessImplTest{},
	}

	for _, gror := range testCaseUpdateSuccess {

		raw, err := ioutil.ReadFile(gror.Url)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &rootobject)

		_, _, actualErr := r.GetList()

		if gror.expectErr != actualErr {
			panic("test case failed for " + gror.Name)
		}
	}
}
