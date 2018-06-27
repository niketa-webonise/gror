package services

import (
	"encoding/json"
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
type CreateFailMarshalTest struct {
}

var testCreateDockerFail = []struct {
	URL       string
	Message   string
	Name      string
	expectErr error
}{
	{
		URL:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data but fails in saving data to db",
		expectErr: errors.New("Failed to save in db"),
	},
}

var testCreateDockerSuccess = []struct {
	URL       string
	Message   string
	Name      string
	expectErr error
}{
	{
		URL:       "../sample_gror.json",
		Message:   "successfully unmarshall and saved in db",
		Name:      "valid data and successfully saved in db",
		expectErr: nil,
	},
}

var testGetitemFail = []struct {
	URL       string
	Message   string
	Name      string
	expectErr error
}{
	{
		URL:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data but fails to get the item",
		expectErr: errors.New("Failed to get the item"),
	},
}

var testGetitemSuccess = []struct {
	URL       string
	Message   string
	Name      string
	expectErr error
}{
	{
		URL:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data and get the item",
		expectErr: nil,
	},
}

var testGetListSuccess = []struct {
	URL       string
	Message   string
	Name      string
	expectErr error
}{
	{
		URL:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data and get the whole list",
		expectErr: nil,
	},
}
var testUpdateFail = []struct {
	URL       string
	Message   string
	Name      string
	expectErr error
}{
	{
		URL:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data but fail to update data",
		expectErr: errors.New("Failed to update data"),
	},
}
var testUpdateSuccess = []struct {
	URL       string
	Message   string
	Name      string
	expectErr error
}{
	{
		URL:       "../sample_gror.json",
		Message:   "successfully unmarshall",
		Name:      "valid data and data updated",
		expectErr: nil,
	},
}

var testCreateFailMarshal = []struct {
	URL       string
	Message   string
	Name      string
	expectErr error
}{
	{
		URL:       "../sample_gror1.json",
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
	s := &InsertDataImpl{
		InsertDockerDaoImpl: CreateFailImplTest{},
	}

	for _, test := range testCreateDockerFail {
		jsonFile, err := ioutil.ReadFile(test.URL)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		actualErr := s.InsertData(jsonFile)
		if test.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", test.expectErr.Error(), actualErr.Error())
		}
	}

	s = &InsertDataImpl{
		InsertDockerDaoImpl: CreateSuccessImplTest{},
	}

	for _, test := range testCreateDockerSuccess {
		jsonFile, err := ioutil.ReadFile(test.URL)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		actualErr := s.InsertData(jsonFile)
		if test.expectErr != actualErr {
			t.Errorf("got code %q but expected %q", test.expectErr.Error(), actualErr.Error())
		}
	}

	s = &InsertDataImpl{
		InsertDockerDaoImpl: CreateFailMarshalTest{},
	}

	for _, test := range testCreateFailMarshal {
		jsonFile, err := ioutil.ReadFile(test.URL)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		actualErr := s.InsertData(jsonFile)
		if test.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", test.expectErr.Error(), actualErr.Error())
		}
	}

}

func TestGetItem(t *testing.T) {
	s := &GetItemImpl{
		GetDockerDaoImpl: GetItemFailImplTest{},
	}

	for _, test := range testGetitemFail {
		raw, err := ioutil.ReadFile(test.URL)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		_, actualErr := s.GetItem(raw)
		if test.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", test.expectErr.Error(), actualErr.Error())
		}
	}

	s = &GetItemImpl{
		GetDockerDaoImpl: GetItemSuccessImplTest{},
	}

	for _, test := range testGetListSuccess {
		josnFile, err := ioutil.ReadFile(test.URL)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		_, actualErr := s.GetItem(josnFile)

		if test.expectErr != actualErr {
			t.Errorf("got code %q but expected %q", test.expectErr.Error(), actualErr.Error())
		}
	}

	s = &GetItemImpl{
		GetDockerDaoImpl: CreateFailMarshalTest{},
	}

	for _, test := range testCreateFailMarshal {
		jsonFile, err := ioutil.ReadFile(test.URL)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		_, actualErr := s.GetItem(jsonFile)
		if test.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", test.expectErr.Error(), actualErr.Error())
		}
	}
}

func TestUpdateData(t *testing.T) {
	var rootobject models.Root
	s := &UpdateDataImpl{
		UpdateDockerDaoImpl: UpdateFailImplTest{},
	}

	for _, test := range testUpdateFail {
		jsonFile, err := ioutil.ReadFile(test.URL)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		json.Unmarshal(jsonFile, &rootobject)

		actualErr := s.UpdateData(jsonFile)
		if test.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", test.expectErr.Error(), actualErr.Error())
		}
	}

	s = &UpdateDataImpl{
		UpdateDockerDaoImpl: UpdateSuccessImplTest{},
	}

	for _, test := range testUpdateSuccess {
		jsonFile, err := ioutil.ReadFile(test.URL)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		json.Unmarshal(jsonFile, &rootobject)
		actualErr := s.UpdateData(jsonFile)
		if test.expectErr != actualErr {
			t.Errorf("got code %q but expected %q", test.expectErr.Error(), actualErr.Error())
		}
	}

	s = &UpdateDataImpl{
		UpdateDockerDaoImpl: CreateFailMarshalTest{},
	}

	for _, test := range testCreateFailMarshal {
		jsonFile, err := ioutil.ReadFile(test.URL)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		json.Unmarshal(jsonFile, &rootobject)
		actualErr := s.UpdateData(jsonFile)
		if test.expectErr.Error() != actualErr.Error() {
			t.Errorf("got code %q but expected %q", test.expectErr.Error(), actualErr.Error())
		}
	}
}
func TestGetList(t *testing.T) {
	var rootobject models.Root
	s := &GetListImpl{
		GetListDockerDaoImpl: GetListSuccessImplTest{},
	}

	for _, test := range testUpdateSuccess {
		jsonFile, err := ioutil.ReadFile(test.URL)
		if err != nil {
			t.Fatal(err)
			os.Exit(1)
		}
		json.Unmarshal(jsonFile, &rootobject)
		_, _, actualErr := s.GetList()
		if test.expectErr != actualErr {
			t.Errorf("got code %q but expected %q", test.expectErr.Error(), actualErr.Error())
		}
	}
}
