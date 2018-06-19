#Gror

Currently, Gror provides the following functionality:

Provides the different API request like,

POST: Accept data through form, Converting forms data to json and bson format and stores it to mongoDB database.

GET: Fetch all data according to the project name

PUT: Accept data through form, Converting forms data to json and bson format and update all details of the project 


##Getting Started

####Prerequisites

First you need to install gopkg.in/mgo.v2 to setup mongoDB drivers and gorilla/mux to set up your router

#####Installing prerequisites

Before installing below mongoDB driver, your system should have installation of [mongoDB](https://www.howtoforge.com/tutorial/install-mongodb-on-ubuntu/).

```
go get gopkg.in/mgo.v2

go get -u github.com/gorilla/mux
```

##Examples

Define the database name by intializing the Dbconfig struct with Dial string and database name 

```golang
    // dbconfig intialize the mongoDB dial and database name
	dbconfig := &database.DbConfig{
		Dial:   "mongodb://127.0.0.1:27017/",
		DbName: "dockerDB",
	}

	// db intialize the database
	db, err := dbconfig.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
```

Intialize the server with database, router, controller, services and model.  

```golang
	// sr intialize the DockerServer
	sr := &servers.DockerServer{
		Db:     db,
		Router: mux.NewRouter(),
		DockerController: &controllers.DockerControllerImpl{
			DockerService: &services.DockerServiceImpl{
				DockerDaoImpl: &models.DockerDaoImpl{
					DB: db,
				},
			},
		},
	}
	//r assigns the server to the RouteWrapper
	r := &routes.RouteWrapper{
		Server: sr,
	}

	r.CreateRoute()
	r.Server.Router.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	err = http.ListenAndServe(":9090", r.Server.Router)
	if err != nil {
		log.Fatal(err)
		return
    }
```

Now define all the routes with their respective controller function 

```golang
// CreateRoute defines the all routes of docker
func (s *RouteWrapper) CreateRoute() {
	s.Server.Router.HandleFunc("/", s.Server.DockerController.DockerForm())
	s.Server.Router.HandleFunc("/docker/config/new", s.Server.DockerController.CreateDockerConfig()).Methods("POST")
	s.Server.Router.HandleFunc("/docker", s.Server.DockerController.DockerListForm())
	s.Server.Router.HandleFunc("/docker/config/list", s.Server.DockerController.GetDockerConfigList()).Methods("GET")
	s.Server.Router.HandleFunc("/docker/config/{id}", s.Server.DockerController.GetDockerConfig()).Methods("GET")
	s.Server.Router.HandleFunc("/docker/config/{id}", s.Server.DockerController.UpdateDockerConfig()).Methods("PUT")

}
```
####Different API request

> Insert data(POST): localhost:9090

> Get all list of projects: localhost:9090/docker/config/list

> Get single project details: localhost:9090/docker/config/{id}

> Update single project details: localhost:9090/docker/config/{id}