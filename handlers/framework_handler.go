package handlers

import (
	"fmt"
	"octopus/config"
	"octopus/models"

	"github.com/gofiber/fiber/v2"
)

var err error

func GetFrameworks(c *fiber.Ctx) error {
	// Using var keyword
	//fmt.Println(c)
	config.Connect()

	// c.Send(config.Select_frameworks())
	return c.Status(200).JSON(config.Select_frameworks())
}

func CreateFramework(c *fiber.Ctx) error {

	f := new(models.Framework)

	if err := c.BodyParser(f); err != nil {
		return err
	}
	fmt.Println(f)
	config.Connect()
	config.Insert_framework(f)

	return c.Status(200).JSON(f)
}
