package main

import (
	"Short_URL/common/config"
	"Short_URL/common/models"
	"Short_URL/server"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	
	models.Setup(c.DBUrl)

	server.SetupAndListen()

	app.Listen(c.Port)

}