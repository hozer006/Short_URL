package main

import (
	"Short_URL/common/config"
	"Short_URL/common/models"
	"Short_URL/server"
	"log"
)

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}


	
	models.Setup(c.DBUrl)

	server.SetupAndListen()


}