package server

import (
	"Short_URL/common/models"
	"Short_URL/utils"
	
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
)

func redirect(c *fiber.Ctx) error {
	shorturl := c.Params("redirect")
	short, err := models.FindByShortUrl(shorturl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find short in DB " + err.Error(),
		}) 
	}
	short.Clicked += 1
	err = models.UpdateShort(short)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}


	return c.Redirect(short.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllShort(c *fiber.Ctx) error {
	short, err := models.GetAllShort()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error getting all short links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(short)

}

func getShort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not parse id " + err.Error(),
		})
	}

	short, err := models.GetShort(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not retreive short from db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(short)
}

func createShort(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var short models.ShortURL
	err := c.BodyParser(&short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if short.Random {
		short.ShortURL = utils.RandomURL(8)
	}

	err = models.CreateShort(short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not create short in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(short)

}

func updateShort(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var short models.ShortURL

	err := c.BodyParser(&short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not parse json " + err.Error(),
		})
	}

	err = models.UpdateShort(short)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not update short link in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(short)
}

func deleteShort(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = models.DeleteShort(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map {
		"message": "short deleted.",
	})
}





func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)

	router.Get("short", getAllShort)
	router.Get("/short/:id", getShort)
	router.Post("/short", createShort)
	router.Patch("/short", updateShort)
	router.Delete("short/:id",deleteShort)
	

	router.Listen(":3000")
}
