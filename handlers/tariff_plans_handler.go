package handlers

import (
	"octopus/config"

	"github.com/gofiber/fiber/v2"
)

func ListPlans(c *fiber.Ctx) error {

	return c.Status(200).JSON(config.Select_tariff_plans())
}
