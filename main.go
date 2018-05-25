package main

import (
	"log"

	"github.com/gror/database"
	"github.com/gror/routes"
)

func main() {

	err := database.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	routes.CreateRoute()
}
