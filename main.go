package main

import (
	"fmt"
	"octopus/candidate"
	"octopus/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	//app.Get("/", helloWorld)

	// framework routes
	app.Get("/api/v1/frameworks", handlers.GetFrameworks)
	app.Post("/api/v1/frameworks", handlers.CreateFramework)
	//
	// Audit activities routes
	app.Post("/api/v1/auditActivities", handlers.InitiateAuditActivity)
	app.Get("/api/v1/auditActivities", handlers.GetAuditActivities)

	app.Get("/api/v1/candidate/:id", candidate.GetCandidate)
	app.Delete("/api/v1/candidate/:id", candidate.DeleteCandidate)

	//SCF Domains routes
	//app.Post("/api/v1/SCFDomains", handlers.InsertDomain)
	app.Post("/api/v1/segments", handlers.PrintingSegment)
	app.Post("/api/v1/GetDomainsFromFrameworks", handlers.GetDomainsFromFrameworks)

	//testing endpoint
	app.Get("/api/v1/testingportion", handlers.TestingPortion2)

	//tariff plans endpoint
	app.Get("/api/v1/tariffplans", handlers.ListPlans)

	//users and clients endpoint
	app.Post("/api/v1/clients", handlers.RegisterClient)
	app.Post("/api/v1/users", handlers.GetUsersByCompanyID)
	app.Post("/api/v1/departments", handlers.GetDepartmentsByCompanyID)
	app.Post("/api/v1/departments/users", handlers.GetUsersByDepartment)

	//SCF Base Controls routes
	app.Get("/api/v1/BaseControls", handlers.GetBaseControls)
	app.Post("/api/v1/BaseControlsPerDomain", handlers.GetBaseControlsPerDom)
	app.Post("/api/v1/ControlDetails", handlers.GetControlDetails)
	app.Post("/api/v1/DeployControls", handlers.DeployControls)
	app.Post("/api/v1/GetEvidenceRequests", handlers.GetEvidenceRequests)
	app.Post("/api/v1/GetEvidenceRequests", handlers.GetEvidenceRequests)
	app.Post("/api/v1/GetEvidenceRequestControls", handlers.GetEvidenceRequestControls)
	//
	app.Get("/api/v1/GetControlsWithDetailsPerFilter/:word", handlers.GetControlsWithDetailsPerFilter)
	app.Post("/api/v1/GetControlsPerFramework", handlers.GetControlsPerFramework)
	//MigrateControlDetails
}

func main() {
	//

	//
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString("Welcone to my go_______ app-++-!")
	})

	setupRoutes(app)

	app.Listen(":3000")
}
