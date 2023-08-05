package candidate

import (
	"github.com/gofiber/fiber/v2"
)

func GetCandidates(c *fiber.Ctx) {
	c.Send("All Candidates")
}

func GetCandidate(c *fiber.Ctx) {
	c.Send("Single Candidate")
}

func NewCandidate(c *fiber.Ctx) {
	c.Send("New Candidate")
}

func DeleteCandidate(c *fiber.Ctx) {
	c.Send("Delete Candidate")
}
