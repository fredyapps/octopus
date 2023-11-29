package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"octopus/config"
	"octopus/models"
	"os"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func convType(v interface{}) interface{} {

	aaaa := v.([]interface{})

	fmt.Println(aaaa[0])

	return aaaa[0]
}

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
		scfcontrol.Scf_domain = convType(result[i]["SCF Control"]).(string)
		scfcontrol.Control_question = convType(result[i]["SCF Control Question"]).(string)
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

	config.Connect()

	for i := 0; i < len(result); i++ {
		//for i := 0; i < 2; i++ {

		var scfcontrol models.SCFcontrol

		scfcontrol.Uuid = fmt.Sprintf("%v", result[i]["UUID"])
		scfcontrol.Scf_control = fmt.Sprintf("%v", convType(result[i]["SCF Control"]))
		scfcontrol.Scf_domain = fmt.Sprintf("%v", result[i]["SCF Domain"])
		scfcontrol.Control_question = fmt.Sprintf("%v", convType(result[i]["SCF Control Question"]))
		scfcontrol.Scf_ref = fmt.Sprintf("%v", result[i]["SCF #"])
		//fmt.Println(scfcontrol)
		config.Insert_control(scfcontrol)

		for index, element := range result[i] {
			//fmt.Println("====================printing control details ================")
			var scfcontroldetail models.SCFcontrolDetail
			scfcontroldetail.Control_uuid = fmt.Sprintf("%v", result[i]["UUID"])
			scfcontroldetail.Control_property = index
			scfcontroldetail.Control_property_value = fmt.Sprintf("%v", element)
			//fmt.Println(scfcontroldetail)
			config.Insert_control_details(scfcontroldetail)

		}

	}

	return c.Status(201).JSON(nil)
}

func Updating_description(c *fiber.Ctx) error {

	jsonFile, err := os.Open("SCF_data.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(byteValue), &result)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/octopus")
	if err != nil {
		fmt.Println(err.Error())
	}

	for i := 0; i < len(result); i++ {
		//for i := 0; i < 2; i++ {
		//fmt.Println(scfcontrol)

		db.Exec("UPDATE scfcontrols SET description = ?  WHERE  uuid = ?", fmt.Sprintf("%v", result[i]["Secure Controls Framework (SCF)\nControl Description"]), fmt.Sprintf("%v", result[i]["UUID"]))

	}
	defer db.Close()
	return c.Status(201).JSON(nil)
}

func Add_BOG_Controls(c *fiber.Ctx) error {

	jsonFile, err := os.Open("BOG_FRAMEWORKS.JSON")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println("====== printing BOG array size ============")
	fmt.Println(len(result))
	fmt.Println("====== done printing BOG array size ============")
	var dsn string = "root:@tcp(127.0.0.1:3306)/octopus"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < len(result); i++ {

		//	contrs := config.Check_if_control_exist(fmt.Sprintf("%v", result[i]["SCF Control"]))
		var controls []string

		db.Select("SCFcontrols.uuid").Table("SCFcontrols").Where("SCFcontrols.scf_control = ? AND SCFcontrols.scf_ref = ?", fmt.Sprintf("%v", result[i]["SCF Control"]), fmt.Sprintf("%v", result[i]["SCF #"])).Find(&controls)
		fmt.Println(controls)

		if len(controls) == 0 {
			var scfcontrol models.SCFcontrol
			uuid := uuid.NewString()
			scfcontrol.Uuid = uuid
			scfcontrol.Scf_control = fmt.Sprintf("%v", result[i]["SCF Control"])
			scfcontrol.Scf_domain = fmt.Sprintf("%v", result[i]["SCF Domain"])
			scfcontrol.Control_question = fmt.Sprintf("%v", result[i]["SCF Control Question"])
			scfcontrol.Scf_ref = fmt.Sprintf("%v", result[i]["SCF #"])
			config.Insert_control(scfcontrol)

			if result[i]["Bank of Ghana CISD"] != nil {
				var scfcontroldetail models.SCFcontrolDetail
				scfcontroldetail.Control_uuid = uuid
				scfcontroldetail.Control_property = "Bank of Ghana CISD"
				scfcontroldetail.Control_property_value = fmt.Sprintf("%v", result[i]["Bank of Ghana CISD"])
				config.Insert_control_details(scfcontroldetail)
			}
			if result[i]["ISO\r\n27002\r\nv2022"] != nil {
				var scfcontroldetail models.SCFcontrolDetail
				scfcontroldetail.Control_uuid = uuid
				scfcontroldetail.Control_property = "ISO\r\n27002\r\nv2022"
				scfcontroldetail.Control_property_value = fmt.Sprintf("%v", result[i]["ISO\r\n27002\r\nv2022"])
				config.Insert_control_details(scfcontroldetail)
			}
			if result[i]["CSA CII"] != nil {
				var scfcontroldetail models.SCFcontrolDetail
				scfcontroldetail.Control_uuid = uuid
				scfcontroldetail.Control_property = "CSA CII"
				scfcontroldetail.Control_property_value = fmt.Sprintf("%v", result[i]["CSA CII"])
				config.Insert_control_details(scfcontroldetail)
			}
			if result[i]["Methods To Comply With SCF Controls"] != nil {
				var scfcontroldetail models.SCFcontrolDetail
				scfcontroldetail.Control_uuid = uuid
				scfcontroldetail.Control_property = "Methods To Comply With SCF Controls"
				scfcontroldetail.Control_property_value = fmt.Sprintf("%v", result[i]["Methods To Comply With SCF Controls"])
				config.Insert_control_details3(scfcontroldetail)
			}
			if result[i]["Relative Control Weighting"] != nil {
				var scfcontroldetail models.SCFcontrolDetail
				scfcontroldetail.Control_uuid = uuid
				scfcontroldetail.Control_property = "Relative Control Weighting"
				scfcontroldetail.Control_property_value = fmt.Sprintf("%v", result[i]["Relative Control Weighting"])
				config.Insert_control_details3(scfcontroldetail)
			}

		} else {

			if result[i]["Bank of Ghana CISD"] != nil {
				var scfcontroldetail models.SCFcontrolDetail
				scfcontroldetail.Control_uuid = controls[0]
				scfcontroldetail.Control_property = "Bank of Ghana CISD"
				scfcontroldetail.Control_property_value = fmt.Sprintf("%v", result[i]["Bank of Ghana CISD"])
				config.Insert_control_details(scfcontroldetail)
			}

			if result[i]["CSA CII"] != nil {
				var scfcontroldetail models.SCFcontrolDetail
				scfcontroldetail.Control_uuid = controls[0]
				scfcontroldetail.Control_property = "CSA CII"
				scfcontroldetail.Control_property_value = fmt.Sprintf("%v", result[i]["CSA CII"])
				config.Insert_control_details(scfcontroldetail)
			}

		}
	}
	return c.Status(201).JSON(nil)
}

func Update_single_control(c *fiber.Ctx) error {

	config.Update_control_with_description("85257ab3-b5c4-47da-a071-82c202b55e38", "this is a desc")

	return c.Status(201).JSON("hhhhhh")
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
