package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"octopus/config"
	"octopus/models"
	"os"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func TestingPortion(c *fiber.Ctx) error {
	fmt.Println("====================printing c body================")
	fmt.Println()
	//var the_field *models.Field
	//_ = json.Unmarshal([]byte(c.Body()), &the_field)

	fmt.Println("====================printing field================")
	//fmt.Println(the_field)

	fmt.Println("====================printing c body type ================")
	fmt.Println(reflect.TypeOf(c.Body()))
	var result []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(c.Body()), &result)

	// Print the data type of result variable
	fmt.Println("====================Print the data type of result variable ================")
	fmt.Println(reflect.TypeOf(result))
	config.Connect()
	for i := 0; i < len(result); i++ {
		fmt.Println("====================looping through json result variable ================")

		var scfcontrol models.SCFcontrol
		scfcontrol.Uuid = fmt.Sprintf("%v", result[i]["UUID"])
		scfcontrol.Scf_control = fmt.Sprintf("%v", result[i]["SCF Domain"])
		scfcontrol.Scf_domain = fmt.Sprintf("%v", result[i]["SCF Control"])
		scfcontrol.Control_question = fmt.Sprintf("%v", result[i]["SCF Control Question"])
		scfcontrol.Scf_ref = fmt.Sprintf("%v", result[i]["SCF #"])
		fmt.Println(scfcontrol)
		config.Insert_control(scfcontrol)

		for index, element := range result[i] {
			fmt.Println("====================printing control details ================")
			var scfcontroldetail models.SCFcontrolDetail
			scfcontroldetail.Control_uuid = fmt.Sprintf("%v", result[i]["UUID"])
			scfcontroldetail.Control_property = index
			scfcontroldetail.Control_property_value = fmt.Sprintf("%v", element)
			fmt.Println(scfcontroldetail)
			config.Insert_control_details(scfcontroldetail)

		}

	}

	return c.Status(200).JSON(result)

}

func TestingPortion2(c *fiber.Ctx) error {

	jsonFile, err := os.Open("SCF_data.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println("====================Printing type of data in file processed ================")
	fmt.Println(reflect.TypeOf(result))

	config.Connect()

	for i := 0; i < len(result); i++ {
		fmt.Println("====================looping through json result variable ================")

		var scfcontrol models.SCFcontrol
		scfcontrol.Uuid = fmt.Sprintf("%v", result[i]["UUID"])
		scfcontrol.Scf_control = fmt.Sprintf("%v", result[i]["SCF Domain"])
		scfcontrol.Scf_domain = fmt.Sprintf("%v", result[i]["SCF Control"])
		scfcontrol.Control_question = fmt.Sprintf("%v", result[i]["SCF Control Question"])
		scfcontrol.Scf_ref = fmt.Sprintf("%v", result[i]["SCF #"])
		fmt.Println(scfcontrol)
		config.Insert_control(scfcontrol)

		for index, element := range result[i] {
			fmt.Println("====================printing control details ================")
			var scfcontroldetail models.SCFcontrolDetail
			scfcontroldetail.Control_uuid = fmt.Sprintf("%v", result[i]["UUID"])
			scfcontroldetail.Control_property = index
			scfcontroldetail.Control_property_value = fmt.Sprintf("%v", element)
			fmt.Println(scfcontroldetail)
			config.Insert_control_details(scfcontroldetail)

		}

	}

	return c.Status(201).JSON(nil)
}

func InsertDomain(c *fiber.Ctx) error {
	// Using var keyword
	var dmn []*models.SCFDomain
	_ = json.Unmarshal([]byte(c.Body()), &dmn)
	// if err := c.BodyParser(dmn); err != nil {
	// 	return err
	// }

	fmt.Println(dmn)
	config.Connect()

	for index, element := range dmn {
		config.Insert_domain(element)
		fmt.Println(index)
	}
	//

	return c.Status(201).JSON(dmn)
}

func PrintingSegment(c *fiber.Ctx) error {

	var segmts *models.Segment

	_ = json.Unmarshal([]byte(c.Body()), &segmts)

	config.Connect()

	for index, element := range segmts.Domains_and_principles.Fields {
		element.Segment = "Domains & Principles"
		config.Insert_schema(element)
		fmt.Println(index)
	}

	for index, element := range segmts.SCF20232.Fields {
		element.Segment = "SCF 2023.2"
		config.Insert_schema(element)
		fmt.Println(index)
	}

	for index, element := range segmts.Assessment_Objectives_20232.Fields {
		element.Segment = "Assessment Objectives 2023.2"
		config.Insert_schema(element)
		fmt.Println(index)
	}

	for index, element := range segmts.Evidence_Request_List_20232.Fields {
		element.Segment = "Evidence Request List 2023.2"
		config.Insert_schema(element)
		fmt.Println(index)
	}

	for index, element := range segmts.Privacy_Management_20232.Fields {
		element.Segment = "Privacy Management 2023.2"
		config.Insert_schema(element)
		fmt.Println(index)
	}

	for index, element := range segmts.Risk_Catalog.Fields {
		element.Segment = "Risk Catalog"
		config.Insert_schema(element)
		fmt.Println(index)
	}

	for index, element := range segmts.Threat_Catalog.Fields {
		element.Segment = "Threat Catalog"
		config.Insert_schema(element)
		fmt.Println(index)
	}

	for index, element := range segmts.Authoritative_Sources.Fields {
		element.Segment = "Authoritative Sources"
		config.Insert_schema(element)
		fmt.Println(index)
	}

	return c.Status(201).JSON(nil)

}
