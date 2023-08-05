package main

import (
	"candidate"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	//app.Get("/", helloWorld)

	app.Get("/api/v1/candidate", candidate.GetCandidates)
	app.Get("/api/v1/candidate/:id", candidate.GetCandidate)
	app.Post("/api/v1/candidate", candidate.NewCandidate)
	app.Delete("/api/v1/candidate/:id", candidate.DeleteCandidate)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	fmt.Println(c.Body())
	// 	return c.SendString("Welcone to my go_______ app-++-!")
	// })

	setupRoutes(app)

	app.Listen(":3000")
}
