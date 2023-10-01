package handlers

import (
	"octopus/config"

	"github.com/gofiber/fiber/v2"
)

func InitiateAuditActivity(c *fiber.Ctx) error {
	// Using var keyword
	//fmt.Println(c)
	config.Connect()

	// c.Send(config.Select_frameworks())
	return c.Status(200).JSON(config.Select_frameworks())
}

func GetAuditActivities(c *fiber.Ctx) error {
	// Using var keyword
	//fmt.Println(c)
	config.Connect()

	// c.Send(config.Select_frameworks())
	return c.Status(200).JSON(config.Select_frameworks())
}
