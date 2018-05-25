package main

import (
	"log"

	"github.com/docker_orchestrator/database"
	"github.com/docker_orchestrator/routes"
)

func main() {

	err := database.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	routes.CreateRoute()
}
