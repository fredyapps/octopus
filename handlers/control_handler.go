package handlers

import (
	"encoding/json"
	"fmt"
	"octopus/config"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func GetBaseControls(c *fiber.Ctx) error {
	// Using var keyword
	//fmt.Println(c)
	config.Connect()
	// config.Select_controls()
	// c.Send(config.Select_frameworks())
	return c.Status(200).JSON(config.Select_controls())
}

func GetBaseControlsPerDom(c *fiber.Ctx) error {
	// Using var keyword
	fmt.Println(reflect.TypeOf(c.Body()))
	var result map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(c.Body()), &result)

	config.Connect()
	return c.Status(200).JSON(config.Select_controls_with_details_per_domain(result["domain"]))
}

func GetControlDetails(c *fiber.Ctx) error {
	// Using var keyword

	fmt.Println(reflect.TypeOf(c.Body()))
	var result map[string]string

	//marshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(c.Body()), &result)

	config.Connect()
	return c.Status(200).JSON(config.Select_control_details(result["uuid"]))
}
