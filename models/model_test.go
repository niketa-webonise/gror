package models

import (
	"errors"
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

// Instance defines all variable of Instance which are to be used to read a json file and bson to interact with mongoDB
type InstanceTest struct {
	EnvMap              EnvMapTest        `json:"EnvMap"` // EnvMap struct type variable which stores the environment map
	PortMapping         string            `json:"PortMapping"`
	AuthId              string            `json:"authId"`
	HostId              string            `json:"hostId"`
	VolumeMapping       VolumeMappingTest `json:"volumeMapping"` // VolumeMapping struct type variable which stores the Volume mapping
	VolumesFrom         string            `json:"volumesFrom"`
	CommandToBeExecuted string            `json:"commandToBeExecuted"`
	Links               string            `json:"links"`
	ImageName           string            `json:"imageName"`
	Tag                 string            `json:"tag"`
	HostsMapping        string            `json:"hostsMapping"`
	Name                string            `json:"name"`
}

// Instance defines all variable of Instance which are to be used to read a json file and bson to interact with mongoDB
type EnvMapTest struct {
	CASSANDRA_BROADCAST_ADDRESS string `json:"CASSANDRA_BROADCAST_ADDRESS"`
	CASSANDRA_SEEDS             string `json:"CASSANDRA_SEEDS"`
}

// VolumeMapping defines all variable of VolumeMapping which are to be used to read a json file and bson to interact with mongoDB
type VolumeMappingTest struct {
	CassData   string `json:"/home/ubuntu/cass-data"`
	CassConfig string `json:"/home/ubuntu/cass-config"`
}

// Component defines all variable of Component which are to be used to read a json file and bson to interact with mongoDB
type ComponentTest struct {
	Instances []InstanceTest `json:"instances"` // Stores the all instances
	Name      string         `json:"name"`
}

// Host defines all variable of Host which are to be used to read a json file and bson to interact with mongoDB
type HostTest struct {
	Protocol                string `json:"protocol"`
	ApiVersion              string `json:"apiVersion"`
	HostType                string `json:"hostType"`
	DockerVersion           string `json:"dockerVersion"`
	Alias                   string `json:"alias"`
	CertPathForDockerDaemon string `json:"certPathForDockerDaemon"`
	IP                      string `json:"ip"`
	DockerPort              string `json:"dockerPort"`
}

// AuthData defines all variable of AuthData which are to be used to read a json file and bson to interact with mongoDB
type AuthDataTest struct {
	Password string `json:"password"`
	Key      string `json:"key"`
	Username string `json:"username"`
	Auth     string `json:"auth"`
	Email    string `json:"email"`
}

// SystemInfo defines all variable of SystemInfo which are to be used to read a json file and bson to interact with mongoDB
type SystemInfoTest struct {
	GrorVersion string `json:"grorVersion"` // Holds the system version
	Name        string `json:"name"`        // Holds the system name
}

// Root defines all variable of Root which are to be used to read a json file and bson to interact with mongoDB
type RootTest struct {
	ID         bson.ObjectId   `json:"id"`         // Holds the unique id of every record
	SystemInfo SystemInfoTest  `json:"systemInfo"` // SystemInfo struct type variable which stores the SystemInfo
	AuthDatas  []AuthDataTest  `json:"authData"`   // Holds the all AuthDatas info
	Hosts      []HostTest      `json:"hosts"`      // Holds the all Hosts info
	Components []ComponentTest `json:"components"` // Holds the all Components info
}

type DockerDaoImplTest struct {
	CreateDockerFunc  func(rootobject Root) error
	GetDockerItemFunc func(rootobject Root) (Root, error)
	GetDockerListFunc func(rootobject Root) ([]string, []string)
	UpdateDockerFunc  func(rootobject Root) error
}

var testCases = []struct {
	Message           string
	expectErr         bool
	DockerDaoImplTest DockerDaoImplTest
}{
	{

		Message:   "successfully tested no errors",
		expectErr: false,
		DockerDaoImplTest: DockerDaoImplTest{
			CreateDockerFunc: func(rootobject Root) error {
				return nil
			},
			GetDockerItemFunc: func(rootobject Root) (Root, error) {
				return rootobject, nil
			},
			GetDockerListFunc: func(rootobject Root) ([]string, []string) {
				return nil, nil
			},
			UpdateDockerFunc: func(rootobject Root) error {
				return nil
			},
		},
	},
	{

		Message:   "unsuccessfull tested with errors",
		expectErr: true,
		DockerDaoImplTest: DockerDaoImplTest{
			CreateDockerFunc: func(rootobject Root) error {
				return errors.New("something went wrong")
			},
			GetDockerItemFunc: func(rootobject Root) (Root, error) {
				return rootobject, errors.New("something went wrong")
			},
			GetDockerListFunc: func(rootobject Root) ([]string, []string) {
				return nil, nil
			},
			UpdateDockerFunc: func(rootobject Root) error {
				return errors.New("something went wrong")
			},
		},
	},
}

// var root = []RootTest{}

func TestCreateDocker(t *testing.T) {

	for _, gror := range testCases {

		fmt.Println(gror.DockerDaoImplTest.CreateDockerFunc)

		if gror.DockerDaoImplTest.CreateDockerFunc != nil && gror.expectErr == false {
			t.Errorf("error= %q, want %q", gror.DockerDaoImplTest.CreateDockerFunc, gror.Message)
		}
	}
}
func TestGetDockerItem(t *testing.T) {

	for _, gror := range testCases {

		if gror.DockerDaoImplTest.GetDockerItemFunc != nil && gror.expectErr == false {
			t.Errorf("error= %q, want %q", gror.DockerDaoImplTest.CreateDockerFunc, gror.Message)
		}
	}
}
func TestUpdateDocker(t *testing.T) {

	for _, gror := range testCases {

		if gror.DockerDaoImplTest.UpdateDockerFunc != nil && gror.expectErr == false {
			t.Errorf("error= %q, want %q", gror.DockerDaoImplTest.CreateDockerFunc, gror.Message)
		}
	}
}
func TestGetDockerList(t *testing.T) {

	for _, gror := range testCases {

		if gror.DockerDaoImplTest.GetDockerListFunc != nil && gror.expectErr == false {
			t.Errorf("error= %q, want %q", gror.DockerDaoImplTest.CreateDockerFunc, gror.Message)
		}
	}
}
func (s DockerDaoImplTest) CreateDocker(rootobject Root) error {
	return nil
}

func (s DockerDaoImplTest) CreateDockerReturns(rootobject Root) error {
	return errors.New("error")
}

func (s DockerDaoImplTest) GetDockerItem(rootobject Root) (Root, error) {
	return rootobject, nil
}
func (s DockerDaoImplTest) UpdateDocker(rootobject Root) error {
	return nil
}
func (s DockerDaoImplTest) GetDockerList(rootobject Root) ([]string, []string) {
	return nil, nil
}
