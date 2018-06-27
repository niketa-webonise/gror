# Docker Configuartion API
- Docker Configuration API implements the GET,POST,PUT RESTAPI.
- The Project uses mongoDB backend.

###### It provides these set of APIs -
	/docker/config/new - (GET request)By opening the URL in browser, the new form will open and user can add the respective fields .
	/docker/config - (POST request)After the user click on submit button, the data will get inserted in mongodb.
	/docker/config- (GET request)By opening the URL in browser, the user able to view the list of docker project names.
	/docker/config/{id} - (GET request)After clicking on one of the docker project names, the user will be able to view the particular docker project 						      details.
	/docker/config/{id} - (PUT request)By clicking on update button,the particular project with updated details will get updated in mongo database.


## Prerequisites-
	1. For routing  - get this package "go get github.com/gorilla/mux"
	2. For using mongodb - get these packages "gopkg.in/mgo.v2/bson" and "gopkg.in/mgo.v2"
	

## For Installing MongoDB follow this Link - 
[https://docs.mongodb.com/manual/installation](https://docs.mongodb.com/manual/installation)

##To Run Test cases-
 	There are two ways. The easy one is to use the -run flag and provide a pattern matching names of the tests you want to run.
 Example:
 ```
 $ go test -run NameOfTest
 ```

 See the [docs](https://golang.org/pkg/testing/#hdr-Subtests_and_Sub_benchmarks) for more info.

 The other way is to name the specific file, containing the tests you want to run:
 ```
 $ go test NameOfTestFile
 ```
 Example:
 ```
 $go test controllers_test.go
 ```

 But there's a catch. This works well if:

 - controller.go is in package controllers.
 - controller_test.go is in package controller_test and imports 'controllers'.
 If controller_test and controller.go is in the same package (a common case) then you must name all other files required to build controller_test.
 In this example it  would  be:
 ```
 $ go test controller_test.go controller.go
 ```

 To run all test files in one go,use this command:

 ```
 $go test ./...
 ```

 If you want to see the what's logged when testing it's worth mentioning the -v (verbose) flag. From the docs -v Verbose output: log all tests as they are run. Also print all text from Log and Logf calls even if the test succeeds.

 Example:
 ```
 $go test -v
 ```




		 
	
	

