package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"octopus/config"
	"octopus/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func DeployControls(c *fiber.Ctx) error {
	// Using var keyword
	fmt.Println(c)
	config.Connect()

	f := new(models.EvidenceRequest)

	if err := c.BodyParser(f); err != nil {
		return err
	}
	f.Req_reference = generateRandomString(10)
	f.Contributors = fmt.Sprintf("%v", f.Contributors)

	config.Insert_evidence_request(f)
	//store deployed controls
	dep_controls := f.Controls.([]interface{})
	for i := 0; i < len(dep_controls); i++ {
		config.Insert_deployed_control(f.Req_reference, fmt.Sprintf("%v", dep_controls[i]))
	}

	rep := new(models.ResponsePayload)
	rep.RESPONSECODE = 201
	rep.RESPONSEMESSAGE = "Controls successfully deployed with evidence request!"

	return c.Status(200).JSON(rep)
}

func GetEvidenceRequests(c *fiber.Ctx) error {

	var req map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(c.Body()), &req)

	return c.Status(200).JSON(config.Select_evidence_requests(req["company"]))
}

func GetEvidenceRequestControls(c *fiber.Ctx) error {

	var req map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(c.Body()), &req)

	return c.Status(200).JSON(config.Select_evidence_request_controls(req["reference"]))
}

func generateRandomString(length int) string {
	result := ""
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	charactersLength := len(characters)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result += string(characters[rand.Intn(charactersLength)])
	}
	return result
}
