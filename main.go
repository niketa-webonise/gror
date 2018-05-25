package main

import (
	"log"

	"github.com/niketa/docker_orchestrator/database"
	"github.com/niketa/docker_orchestrator/routes"
)

func main() {

	err := database.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	routes.CreateRoute()
}
