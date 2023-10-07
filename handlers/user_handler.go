package handlers

import (
	"encoding/json"
	"octopus/config"
	"octopus/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsersByCompanyID(c *fiber.Ctx) error {

	var req map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(c.Body()), &req)

	return c.Status(200).JSON(config.Select_users_per_client(req["company"]))
}

func CreateUser(c *fiber.Ctx) error {

	return c.SendString("User has been created")
}

func RegisterClient(c *fiber.Ctx) error {

	clt := new(models.Client)

	if err := c.BodyParser(clt); err != nil {
		return err
	}
	//fmt.Println(rand.String(10))
	config.Connect()
	config.Insert_client(clt)

	return c.Status(200).JSON(clt)

}

func GetDepartmentsByCompanyID(c *fiber.Ctx) error {

	var req map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(c.Body()), &req)
	//Select_department_per_client

	return c.Status(200).JSON(config.Select_department_per_client(req["company"]))
}

func GetUsersByDepartment(c *fiber.Ctx) error {

	var req map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(c.Body()), &req)
	//Select_department_per_client

	return c.Status(200).JSON(config.Select_users_per_department(req["company"], req["department"]))
}
