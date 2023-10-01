package candidate

import (
	"fmt"
	"octopus/config"
	"octopus/models"

	"github.com/gofiber/fiber/v2"
)

var err error

func GetCandidate(c *fiber.Ctx) error {
	return c.SendString("Single Candidate")
}

func NewCandidate(c *fiber.Ctx) error {

	//var db
	//var MYDB *sql.DB
	//fmt.Println(c.Context())

	candidate := new(models.Candidate)

	if err := c.BodyParser(candidate); err != nil {
		fmt.Println(err)
		return c.Status(503).Send([]byte(err.Error()))

	}

	//config.Database.Create(&candidate)
	//fmt.Println(candidate)
	config.Connect()
	config.Insert_candidate()

	return c.Status(201).JSON(candidate)

}

func DeleteCandidate(c *fiber.Ctx) error {
	return c.SendString("Delete Candidate")
}

// npm install -g nodemon
// sudo npm install -g nodemon
// nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go
