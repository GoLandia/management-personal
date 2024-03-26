package main

import (
	"management-personal/api"
	"management-personal/config"
	"management-personal/database"
	"management-personal/utils"
)

func main() {
	var err error

	err = config.StartConfig()
	utils.FatalError(err)

	d := database.Database{}
	err = database.Connect(&d)
	utils.FatalError(err)

	service := api.NewService()
	service.Start()
}
