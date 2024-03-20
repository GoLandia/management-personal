package main

import (
	"management-personal/database"
	"management-personal/routes"
)

func main() {
	database.Connect()
	routes.Api()
}
