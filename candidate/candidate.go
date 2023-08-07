package octopus

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetCandidates(c *fiber.Ctx) {

	// Using var keyword
	var my_value_1 string
	my_value_1 = "GeeksforGeeks"
	fmt.Println(my_value_1)
	c.SendString("Hello, World 👋!")
}

func GetCandidate(c *fiber.Ctx) {
	c.SendString("Single Candidate")
}

func NewCandidate(c *fiber.Ctx) {
	c.SendString("New Candidate")
}

func DeleteCandidate(c *fiber.Ctx) {
	c.SendString("Delete Candidate")
}
