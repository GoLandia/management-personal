package main

import (
	"log"
	"management-personal/config"
	"management-personal/database"
	"management-personal/routes"
)

func main() {
	if err := config.StartConfig(); err != nil {
		log.Fatal(err)
	}

	d := database.Database{}
	database.Connect(&d)

	routes.Api()
}
