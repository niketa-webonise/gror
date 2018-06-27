package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Instance defines all fields with json and bson field tag
type Instance struct {
	EnvMap              EnvMap        `json:"envMap" bson:"envMap"` //EnvMap field is of type EnvMap
	PortMapping         string        `json:"portMapping" bson:"portMapping"`
	AuthID              string        `json:"authId" bson:"authId"`
	HostID              string        `json:"hostId" bson:"hostId"`
	VolumeMapping       VolumeMapping `json:"volumeMapping" bson:"volumeMapping"` //VolumeMapping field is of type VolumeMapping
	VolumesFrom         string        `json:"volumesFrom" bson:"volumesFrom"`
	CommandToBeExecuted string        `json:"commandToBeExecuted" bson:"commandToBeExecuted"`
	Links               string        `json:"links" bson:"links"`
	ImageName           string        `json:"imageName" bson:"imageName"`
	Tag                 string        `json:"tag" bson:"tag"`
	HostsMapping        string        `json:"hostsMapping" bson:"hostsMapping"`
	Name                string        `json:"name" bson:"name"`
}

//EnvMap defines all fields with json and bson field tag
type EnvMap struct {
	CASSANDRA_BROADCAST_ADDRESS string `json:"CASSANDRA_BROADCAST_ADDRESS" bson:"CASSANDRA_BROADCAST_ADDRESS"`
	CASSANDRA_SEEDS             string `json:"CASSANDRA_SEEDS" bson:"CASSANDRA_SEEDS"`
}

//VolumeMapping defines all fields with json and bson field tag
type VolumeMapping struct {
	CassData   string `json:"/home/ubuntu/cass-data" bson:"/home/ubuntu/cass-data"`
	CassConfig string `json:"/home/ubuntu/cass-config" bson:"/home/ubuntu/cass-config"`
}

//Component defines all fields with json and bson field tag
type Component struct {
	Instances []Instance `json:"instances" bson:"instances"` //Instances field is type of Instance array
	Name      string     `json:"name" bson:"name"`
}

//Host defines all fields with json and bson field tag
type Host struct {
	Protocol                string `json:"protocol" bson:"protocol"`
	ApIVersion              string `json:"apiVersion" bson:"apiVersion"`
	HostType                string `json:"hostType" bson:"hostType"`
	DockerVersion           string `json:"dockerVersion" bson:"dockerVersion"`
	Alias                   string `json:"alias" bson:"alias"`
	CertPathForDockerDaemon string `json:"certPathForDockerDaemon" bson:"certPathForDockerDaemon"`
	IP                      string `json:"ip" bson:"ip"`
	DockerPort              string `json:"dockerPort" bson:"dockerPort"`
}

//AuthData defines all fields with json and bson field tag
type AuthData struct {
	Password string `json:"password" bson:"password"`
	Key      string `json:"key" bson:"key"`
	Username string `json:"username" bson:"username"`
	Auth     string `json:"auth" bson:"auth"`
	Email    string `json:"email" bson:"email"`
}

//SystemInfo defines all fields with json and bson field tag
type SystemInfo struct {
	GrorVersion string `json:"grorVersion" bson:"grorVersion"`
	Name        string `json:"name" bson:"name"`
}

//Root is outermost container that contains SystemInfo
type Root struct {
	ID         bson.ObjectId `json:"id" bson:"_id",omitempty`
	SystemInfo SystemInfo    `json:"systemInfo" bson:"systemInfo"` //SystemInfo field is of type SystemInfo
	AuthDatas  []AuthData    `json:"authData" bson:"authData"`     //AuthDatas field is of type AuthData array
	Hosts      []Host        `json:"hosts" bson:"hosts"`           //Hosts field is of type Host array
	Components []Component   `json:"components" bson:"components"` //Components field is of type array Component
}

type CreateDockerInterface interface {
	CreateDocker(rootobject Root) error
}

type GetDockerItemInterface interface {
	GetDockerItem(rootobject Root) (Root, error)
}

type UpdateDockerInterface interface {
	UpdateDocker(rootobject Root) error
}

type GetDockerListInterface interface {
	GetDockerList() ([]string, []string)
}

//DockerDaoImpl wraps mongoDB Database
type DockerDaoImpl struct {
	DB *mgo.Database
}

//CreateDocker function insert the rootobject of type Root in Database
func (s *DockerDaoImpl) CreateDocker(rootobject Root) error {
	c := s.DB.C("dockers")
	return c.Insert(rootobject)
}

//GetDockerItem function get  rootobject of type Root by id from Database
func (s *DockerDaoImpl) GetDockerItem(rootobject Root) (Root, error) {
	c := s.DB.C("dockers")
	err := c.FindId(rootobject.ID).One(&rootobject)
	return rootobject, err
}

//GetDockerList function returns the id and name from database
func (s *DockerDaoImpl) GetDockerList() ([]string, []string) {
	var rootobject Root
	var names []string
	var ids []string
	c := s.DB.C("dockers")
	find := c.Find(bson.M{})
	items := find.Iter()
	for items.Next(&rootobject) {
		names = append(names, rootobject.SystemInfo.Name)
		ids = append(ids, rootobject.ID.Hex())
	}
	return names, ids
}

//UpdateDocker updates the record of specific id
func (s *DockerDaoImpl) UpdateDocker(rootobject Root) error {
	c := s.DB.C("dockers")
	return c.UpdateId(rootobject.ID, &rootobject)
}
