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
	return c.Status(200).JSON(config.Select_controls_with_details_per_domain(result["domain"], result["framework"]))
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

func GetControlsWithDetailsPerFilter(c *fiber.Ctx) error {
	// Using var keyword
	//c.
	config.Connect()
	return c.Status(200).JSON(config.Select_control_join_details(c.Params("word")))
}

func GetControlsPerFramework(c *fiber.Ctx) error {
	// Using var keyword
	fmt.Println(reflect.TypeOf(c.Body()))
	var result map[string]string

	//marshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(c.Body()), &result)
	config.Connect()
	return c.Status(200).JSON(config.Select_all_controls_per_framework(result["framework"]))
}

func GetDomainsFromFrameworks(c *fiber.Ctx) error {

	fmt.Println(reflect.TypeOf(c.Body()))
	var result map[string][]string

	json.Unmarshal([]byte(c.Body()), &result)
	config.Connect()
	fmt.Println("=============Printing the selected frameworks========")
	fmt.Println(result["frameworks"])
	return c.Status(200).JSON(config.Select_domains_of_selected_frameworks(result["frameworks"]))

	//return c.Status(200).JSON("")

}

//Select_all_controls_per_framework
