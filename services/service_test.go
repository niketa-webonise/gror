package services

import (
	"errors"
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

var testCaseCreateFail = []struct {
	Url              string
	Message          string
	Name             string
	expectErr        error
	serviceInterface CreateFailImplTest
}{
	{
		Url:              "../sample_gror.json",
		Message:          "successfully unmarshall",
		Name:             "valid data but fails in saving data to db",
		expectErr:        errors.New("Failed to save in db"),
		serviceInterface: CreateFailImplTest{},
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

var testCaseUnmarshallingFail = []struct {
	Url       string
	Message   string
	Name      string
	expectErr error
}{
	{
		Url:       "../sample2_gror.json",
		Message:   "unsuccessfully unmarshall",
		Name:      "invalid data failed to unmarshalling",
		expectErr: errors.New("json: cannot unmarshal number into Go struct field SystemInfo.name of type string"),
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
	Url              string
	Message          string
	Name             string
	expectErr        error
	serviceInterface GetItemSuccessImplTest
}{
	{
		Url:              "../sample_gror.json",
		Message:          "successfully unmarshall",
		Name:             "valid data and get the item",
		expectErr:        nil,
		serviceInterface: GetItemSuccessImplTest{},
	},
}

var testCaseGetListSuccess = []struct {
	Url              string
	Message          string
	Name             string
	expectErr        error
	serviceInterface GetListSuccessImplTest
}{
	{
		Url:              "../sample_gror.json",
		Message:          "successfully unmarshall",
		Name:             "valid data and get the whole list",
		expectErr:        nil,
		serviceInterface: GetListSuccessImplTest{},
	},
}
var testCaseUpdateFail = []struct {
	Url              string
	Message          string
	Name             string
	expectErr        error
	serviceInterface UpdateFailImplTest
}{
	{
		Url:              "../sample_gror.json",
		Message:          "successfully unmarshall",
		Name:             "valid data but fail to update data",
		expectErr:        errors.New("Failed to update data"),
		serviceInterface: UpdateFailImplTest{},
	},
}
var testCaseUpdateSuccess = []struct {
	Url              string
	Message          string
	Name             string
	expectErr        error
	serviceInterface UpdateSuccessImplTest
}{
	{
		Url:              "../sample_gror.json",
		Message:          "successfully unmarshall",
		Name:             "valid data and data updated",
		expectErr:        nil,
		serviceInterface: UpdateSuccessImplTest{},
	},
}

func (s CreateSuccessImplTest) CreateDocker(rootobject models.Root) error {
	return nil

}
func (s CreateFailImplTest) CreateDocker(rootobject models.Root) error {
	return errors.New("Failed to save in db")

}
func (s GetItemFailImplTest) GetDockerItem(rootobject models.Root) (models.Root, error) {

	return rootobject, errors.New("Failed to get the item")
}
func (s GetItemSuccessImplTest) GetDockerItem(rootobject models.Root) (models.Root, error) {

	return rootobject, nil
}
func (s UpdateFailImplTest) UpdateDocker(rootobject models.Root) error {

	return errors.New("Failed to update data")
}
func (s UpdateSuccessImplTest) UpdateDocker(rootobject models.Root) error {

	return nil
}
func (s GetListSuccessImplTest) GetDockerList(rootobject models.Root) ([]string, []string) {

	return []string{}, []string{}
}

func TestInsertData(t *testing.T) {

	r := &InsertDataDockerServiceImpl{
		CreateDockerDaoImpl: CreateFailImplTest{},
	}

	for _, testCase := range testCaseCreateFail {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}

		actualErr := r.InsertData(raw)
		if testCase.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}
	}
	for _, testCase := range testCaseUnmarshallingFail {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}

		actualErr := r.InsertData(raw)
		if testCase.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}
	}
	r = &InsertDataDockerServiceImpl{
		CreateDockerDaoImpl: CreateSuccessImplTest{},
	}

	for _, testCase := range testCaseCreateSuccess {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}

		actualErr := r.InsertData(raw)
		if testCase.expectErr != actualErr {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}

	}
}

func TestGetItem(t *testing.T) {

	r := &GetItemDockerServiceImpl{
		GetItemDockerDaoImpl: GetItemFailImplTest{},
	}

	for _, testCase := range testCaseGetitemFail {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}

		_, actualErr := r.GetItem(raw)
		if testCase.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}
	}
	for _, testCase := range testCaseUnmarshallingFail {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}

		_, actualErr := r.GetItem(raw)
		if testCase.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}
	}
	r = &GetItemDockerServiceImpl{
		GetItemDockerDaoImpl: GetItemSuccessImplTest{},
	}

	for _, testCase := range testCaseGetitemSuccess {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		_, actualErr := r.GetItem(raw)
		if testCase.expectErr != actualErr {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}
	}
}

func TestGetList(t *testing.T) {

	r := &GetListDockerServiceImpl{
		GetDockerListDaoImpl: GetListSuccessImplTest{},
	}

	for _, testCase := range testCaseGetListSuccess {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		_, _, actualErr := r.GetList(raw)

		if testCase.expectErr != actualErr {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}
	}
	for _, testCase := range testCaseUnmarshallingFail {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		_, _, actualErr := r.GetList(raw)
		if testCase.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}
	}
}

func TestUpdateData(t *testing.T) {

	r := &UpdateDockerServiceImpl{
		UpdateDockerDaoImpl: UpdateFailImplTest{},
	}

	for _, testCase := range testCaseUpdateFail {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		actualErr := r.UpdateData(raw)
		if testCase.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}
	}
	for _, testCase := range testCaseUnmarshallingFail {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}

		actualErr := r.UpdateData(raw)
		if testCase.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}
	}
	r = &UpdateDockerServiceImpl{
		UpdateDockerDaoImpl: UpdateSuccessImplTest{},
	}

	for _, testCase := range testCaseUpdateSuccess {

		raw, err := ioutil.ReadFile(testCase.Url)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}

		actualErr := r.UpdateData(raw)
		if testCase.expectErr != actualErr {
			t.Errorf("got code %q but expected %q", testCase.expectErr.Error(), actualErr.Error())
		}
	}
}
