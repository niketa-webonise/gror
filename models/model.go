package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Instance defines all variable of Instance which are to be used to read a json file and bson to interact with mongoDB
type Instance struct {
	EnvMap              EnvMap        `json:"EnvMap" bson:"EnvMap"` // EnvMap struct type variable which stores the environment map
	PortMapping         string        `json:"PortMapping" bson:"portMapping"`
	AuthId              string        `json:"authId" bson:"authId"`
	HostId              string        `json:"hostId" bson:"hostId"`
	VolumeMapping       VolumeMapping `json:"volumeMapping" bson:"volumeMapping"` // VolumeMapping struct type variable which stores the Volume mapping
	VolumesFrom         string        `json:"volumesFrom" bson:"volumesFrom"`
	CommandToBeExecuted string        `json:"commandToBeExecuted" bson:"commandToBeExecuted"`
	Links               string        `json:"links" bson:"links"`
	ImageName           string        `json:"imageName" bson:"imageName"`
	Tag                 string        `json:"tag" bson:"tag"`
	HostsMapping        string        `json:"hostsMapping" bson:"hostsMapping"`
	Name                string        `json:"name" bson:"name"`
}

// Instance defines all variable of Instance which are to be used to read a json file and bson to interact with mongoDB
type EnvMap struct {
	CASSANDRA_BROADCAST_ADDRESS string `json:"CASSANDRA_BROADCAST_ADDRESS" bson:"CASSANDRA_BROADCAST_ADDRESS"`
	CASSANDRA_SEEDS             string `json:"CASSANDRA_SEEDS" bson:"CASSANDRA_SEEDS"`
}

// VolumeMapping defines all variable of VolumeMapping which are to be used to read a json file and bson to interact with mongoDB
type VolumeMapping struct {
	CassData   string `json:"/home/ubuntu/cass-data" bson:"/home/ubuntu/cass-data"`
	CassConfig string `json:"/home/ubuntu/cass-config" bson:"/home/ubuntu/cass-config"`
}

// Component defines all variable of Component which are to be used to read a json file and bson to interact with mongoDB
type Component struct {
	Instances []Instance `json:"instances" bson:"instances"` // Stores the all instances
	Name      string     `json:"name" bson:"name"`
}

// Host defines all variable of Host which are to be used to read a json file and bson to interact with mongoDB
type Host struct {
	Protocol                string `json:"protocol" bson:"protocol"`
	ApiVersion              string `json:"apiVersion" bson:"apiVersion"`
	HostType                string `json:"hostType" bson:"hostType"`
	DockerVersion           string `json:"dockerVersion" bson:"dockerVersion"`
	Alias                   string `json:"alias" bson:"alias"`
	CertPathForDockerDaemon string `json:"certPathForDockerDaemon" bson:"certPathForDockerDaemon"`
	IP                      string `json:"ip" bson:"ip"`
	DockerPort              string `json:"dockerPort" bson:"dockerPort"`
}

// AuthData defines all variable of AuthData which are to be used to read a json file and bson to interact with mongoDB
type AuthData struct {
	Password string `json:"password" bson:"password"`
	Key      string `json:"key" bson:"key"`
	Username string `json:"username" bson:"username"`
	Auth     string `json:"auth" bson:"auth"`
	Email    string `json:"email" bson:"email"`
}

// SystemInfo defines all variable of SystemInfo which are to be used to read a json file and bson to interact with mongoDB
type SystemInfo struct {
	GrorVersion string `json:"grorVersion" bson:"grorVersion"` // Holds the system version
	Name        string `json:"name" bson:"name"`               // Holds the system name
}

// Root defines all variable of Root which are to be used to read a json file and bson to interact with mongoDB
type Root struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`                // Holds the unique id of every record
	SystemInfo SystemInfo    `json:"systemInfo" bson:"systemInfo"` // SystemInfo struct type variable which stores the SystemInfo
	AuthDatas  []AuthData    `json:"authData" bson:"authData"`     // Holds the all AuthDatas info
	Hosts      []Host        `json:"hosts" bson:"hosts"`           // Holds the all Hosts info
	Components []Component   `json:"components" bson:"components"` // Holds the all Components info
}

// DockerDao wraps the all method which are interact with database
type CreateDockerDao interface {
	CreateDocker(rootobject Root) error
}
type GetDockerItemDao interface {
	GetDockerItem(rootobject Root) (Root, error)
}
type UpdateDockerItemDao interface {
	UpdateDocker(rootobject Root) error
}
type GetDockerListDao interface {
	GetDockerList(rootobject Root) ([]string, []string)
}

// DockerDaoImpl defines the mongoDB database
type DockerDaoImpl struct {
	DB *mgo.Database
}

// CreateDocker insert the rootobject into the database
func (s *DockerDaoImpl) CreateDocker(rootobject Root) error {
	c := s.DB.C("dockers")
	return c.Insert(rootobject)
}

// GetDockerItem returns the single item from the database
func (s *DockerDaoImpl) GetDockerItem(rootobject Root) (Root, error) {
	c := s.DB.C("dockers")
	err := c.FindId(rootobject.ID).One(&rootobject)
	return rootobject, err
}

// GetDockerList returns the names and object id of all system
func (s *DockerDaoImpl) GetDockerList(rootobject Root) ([]string, []string) {
	var names []string
	var objid []string
	c := s.DB.C("dockers")
	find := c.Find(bson.M{})
	items := find.Iter()
	for items.Next(&rootobject) {
		names = append(names, rootobject.SystemInfo.Name)
		objid = append(objid, rootobject.ID.Hex())

	}
	return names, objid
}

// UpdateDocker updates the record of specific id
func (s *DockerDaoImpl) UpdateDocker(rootobject Root) error {
	c := s.DB.C("dockers")
	return c.UpdateId(rootobject.ID, &rootobject)
}
