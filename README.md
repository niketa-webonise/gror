# Docker Configuartion API
- Docker Configuration API implements the GET,POST,PUT RESTAPI.
- The Project uses mongoDB backend.

###### It provides these set of APIs -
	`localhost:8080/docker/config/new` - (GET request)By opening the URL in browser, the new form will open and user can add the respective fields .
	`localhost:8080/docker/config` - (POST request)After the user click on submit button, the data will get inserted in mongodb.
	`localhost:8080/docker/config`- (GET request)By opening the URL in browser, the user able to view the list of docker project names.
	`localhost:8080/docker/config/{id}` - (GET request)After clicking on one of the docker project names, the user will be able to view the particular docker project 						      details.
	`localhost:8080/docker/config/{id}` - (PUT request)By clicking on update button,the particular project with updated details will get updated in mongo database.


## Prerequisites-
	1. For routing  - get this package "go get github.com/gorilla/mux"
	2. For using mongodb - get these packages "gopkg.in/mgo.v2/bson" and "gopkg.in/mgo.v2"
	

## For Installing MongoDB follow this Link - 
[https://docs.mongodb.com/manual/installation](https://docs.mongodb.com/manual/installation)



		 
	
	

