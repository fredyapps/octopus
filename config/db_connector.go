package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"octopus/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect to the database
// var Database *gorm.DB
var err error

//var dsn string = "root:@tcp(127.0.0.1:3306)/octopus"

var dsn string = "root:Mathapelo@2030@tcp(127.0.0.1:3306)/octopus"

var insert sql.Result

func Connect() {

	// Open the database.
	db, err := sql.Open("mysql", dsn)

	fmt.Println("================ printing the DB stats  line 31 =================")
	fmt.Println(db.Stats())
	//fmt.Println(db.Ping())

	//log.Fatalf("impossible to create the connection: %s", err)

	if err != nil {
		fmt.Println("================ Error  connector line 40 =================")
		fmt.Println(err.Error())
		log.Fatalln(err)
	}

	fmt.Println("================ Trace connector line 45 =================")
	//fmt.Println(db)
	defer db.Close()

}

func Update_control_with_description(uuid string, descr string) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	updated, ok := db.Exec("UPDATE scfcontrols SET description = ?  WHERE  uuid = ?", descr, uuid)
	fmt.Println(updated)
	fmt.Println(ok)

}

func Select_domains_of_selected_frameworks(frms []string) []models.SCFDomain {

	var domains []models.SCFDomain
	var the_domains []models.SCFDomain
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	db.Table("scfdomains").Distinct("scfdomains.SCFDomain", "scfdomains.SCFIdentifier").Joins("JOIN scfcontrols ON scfdomains.SCFDomain=scfcontrols.scf_domain").Joins("JOIN scfcontroldetails ON scfcontrols.uuid=scfcontroldetails.control_uuid").Where("scfcontroldetails.control_property IN ?", frms).Find(&domains)

	for key, domain := range domains {

		fmt.Println(key, " = ", domain)
		domain.Controls = Select_controls_from_selected_frameworks(frms, domain.SCFDomain)

		the_domains = append(the_domains, domain)

	}

	return the_domains

}

func Select_frameworks() []models.Framework {

	var frameworks []models.Framework

	q := "SELECT  id_framework, description ,reference , name , version FROM `frameworks`"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var framework models.Framework
		err = results.Scan(&framework.Id_framework, &framework.Description, &framework.Reference, &framework.Name, &framework.Version)
		if err != nil {
			panic(err.Error())
		}

		framework.Number_of_controls = Select_control_count(framework.Reference)
		framework.Number_of_domains = Select_domains_count(framework.Reference)
		frameworks = append(frameworks, framework)

	}
	defer results.Close()

	return frameworks

}

func Insert_framework(frmk *models.Framework) {

	//=========================================== |||||||||||||| ============================================
	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `frameworks` (`name`, `reference`, `version`,`description`) VALUES (?, ?, ?, ?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, frmk.Name, frmk.Reference, frmk.Version, frmk.Description)
	fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)
	if err != nil {
		fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Framework: %s", err)
	}

	defer db.Close()

}

func Insert_candidate() {

	q := "INSERT INTO `contacts_tab` (`fullName`, `subject`, `email`,`message`) VALUES ('Anani sallah', 'insertion', 'fredy@gmail.com', 'holla')"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("================ Error  connector line 62 =================")
		fmt.Println(err.Error())
	}
	defer db.Close()

	fmt.Println("================ ready for insertion line 56 =================")
	// perform a db.Query insert
	insert, err := db.Query(q)

	fmt.Println("================ after insertion line 60 =================")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	//==================================================================
}

// func Insert_domain(dmn *models.SCFDomain) {

// 	//=========================================== |||||||||||||| ============================================
// 	db, err := sql.Open("mysql", dsn)
// 	query := "INSERT INTO `SCFDomains` (`UUID`, `ID`, `SCFDomain`,`SCFIdentifier`,`SecurityPrivacy`,`PrincipleIntent`) VALUES (?, ?, ?, ?, ?, ?)"

// 	insertResult, err := db.ExecContext(context.Background(), query, dmn.UUID, dmn.ID, dmn.SCFDomain, dmn.SCFIdentifier, dmn.SecurityPrivacy[0], dmn.PrincipleIntent)
// 	fmt.Println("================ Error  connector line 88 =================")
// 	fmt.Println(insertResult)

// 	if err != nil {
// 		fmt.Println("================ Error  connector line 91 =================")
// 		log.Fatalf("impossible insert Domain: %s", err)
// 	}

// 	defer db.Close()

// }

func Insert_schema(fld models.Field) {

	//=========================================== |||||||||||||| ============================================
	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `json_schema_fields` (`segment`, `name`, `type`) VALUES (?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, fld.Segment, fld.Name, fld.Type)
	fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)

	if err != nil {
		fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Domain: %s", err)
	}

	defer db.Close()

}

func Insert_control(SCFctrl models.SCFcontrol) {

	//=========================================== |||||||||||||| ============================================
	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `scfcontrols` (`uuid`, `scf_control`, `scf_domain`,`scf_ref`,`control_question`) VALUES (?, ?, ?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, SCFctrl.Uuid, SCFctrl.Scf_control, SCFctrl.Scf_domain, SCFctrl.Scf_ref, SCFctrl.Control_question)
	//fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)

	if err != nil {
		//fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Domain: %s", err)
	}

	defer db.Close()

}

func Insert_control_details(ctrlDet models.SCFcontrolDetail) {

	//=========================================== |||||||||||||| ============================================
	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `scfcontroldetails` (`control_uuid`, `control_property`, `control_property_value`) VALUES (?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, ctrlDet.Control_uuid, ctrlDet.Control_property, ctrlDet.Control_property_value)
	//fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)

	if err != nil {
		//fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Domain: %s", err)
	}

	defer db.Close()

}

func Insert_control_details3(ctrlDet models.SCFcontrolDetail) {

	//=========================================== |||||||||||||| ============================================
	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `scfcontroldetails3` (`control_uuid`, `control_property`, `control_property_value`) VALUES (?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, ctrlDet.Control_uuid, ctrlDet.Control_property, ctrlDet.Control_property_value)
	//fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)

	if err != nil {
		//fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Domain: %s", err)
	}

	defer db.Close()

}

func Check_if_control_exist(control string) []string {

	var controls []string
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Select("scfcontrols.uuid").Table("scfcontrols").Where("scfcontrols.scf_control = ?", control).Find(&controls)

	return controls

}

func Insert_client(octo_clt *models.Client) {

	//=========================================== |||||||||||||| ============================================
	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `clients` (`company_name`, `country`, `contact_person`,`contact_email`,`tariff_plan`) VALUES (?, ?, ?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, octo_clt.Company_name, octo_clt.Country, octo_clt.Contact_person, octo_clt.Contact_email, octo_clt.Tariff_plan)
	fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)

	if err != nil {
		fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Domain: %s", err)
	}

	defer db.Close()

}

func Select_controls() []models.SCFcontrol {

	var controls []models.SCFcontrol

	q := "SELECT  uuid, scf_control ,scf_domain , scf_ref , control_question FROM `scfcontrols`"

	//
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var control models.SCFcontrol
		err = results.Scan(&control.Uuid, &control.Scf_control, &control.Scf_domain, &control.Scf_ref, &control.Control_question)
		if err != nil {
			panic(err.Error())
		}
		controls = append(controls, control)

	}
	defer results.Close()

	return controls

}

func Select_controls_per_domain(domain string) []models.SCFcontrol {

	var controls []models.SCFcontrol

	q := "SELECT  uuid, scf_control ,scf_domain , scf_ref , control_question FROM `scfcontrols` WHERE  scf_domain ='" + domain + "'"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var control models.SCFcontrol
		err = results.Scan(&control.Uuid, &control.Scf_control, &control.Scf_domain, &control.Scf_ref, &control.Control_question)
		if err != nil {
			panic(err.Error())
		}
		controls = append(controls, control)

	}
	defer results.Close()

	return controls

}

func Select_users_per_client(comp_id string) []models.OctopusUser {

	var employees []models.OctopusUser

	q := "SELECT  firstname, lastname ,email , department , position , user_role , on_leave  FROM `users` WHERE  company_id ='" + comp_id + "'"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var usr models.OctopusUser
		err = results.Scan(&usr.Firstname, &usr.Lastname, &usr.Email, &usr.Department, &usr.Position, &usr.User_role, &usr.On_leave)
		if err != nil {
			panic(err.Error())
		}
		employees = append(employees, usr)

	}

	fmt.Println(employees)
	defer results.Close()

	return employees

}

func Select_control_details(uuid string) []models.SCFcontrolDetail {

	var ctrl_details []models.SCFcontrolDetail
	q := "SELECT control_uuid ,control_property , control_property_value  FROM `scfcontroldetails` WHERE  control_uuid ='" + uuid + "'"
	db, err := sql.Open("mysql", dsn)

	results, err := db.Query(q)

	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	for results.Next() {

		var ctrl_detail models.SCFcontrolDetail
		err = results.Scan(&ctrl_detail.Control_uuid, &ctrl_detail.Control_property, &ctrl_detail.Control_property_value)

		ctrl_details = append(ctrl_details, ctrl_detail)
	}

	defer results.Close()

	return ctrl_details

}

func Select_control_details_with_filter(uuid string, word string) []models.SCFcontrolDetail {

	var ctrl_details []models.SCFcontrolDetail
	q := "SELECT control_uuid ,control_property , control_property_value  FROM `scfcontroldetails` WHERE  control_uuid ='" + uuid + "' AND control_property LIKE '%" + word + "%'"
	db, err := sql.Open("mysql", dsn)

	results, err := db.Query(q)

	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	for results.Next() {

		var ctrl_detail models.SCFcontrolDetail
		err = results.Scan(&ctrl_detail.Control_uuid, &ctrl_detail.Control_property, &ctrl_detail.Control_property_value)

		ctrl_details = append(ctrl_details, ctrl_detail)
	}

	defer results.Close()

	return ctrl_details

}

func Select_controls_with_details_per_domain(domain string, framework string) []models.SCFcontrol {

	var controls []models.SCFcontrol

	//q := "SELECT  uuid, scf_control ,scf_domain , scf_ref , control_question FROM `SCFcontrols` WHERE  scf_domain ='" + domain + "'"
	q := "SELECT scfcontrols.uuid, scfcontrols.scf_control, scfcontrols.scf_domain, scfcontrols.scf_ref, scfcontrols.control_question, scfcontroldetails.control_property,scfcontroldetails.control_property_value FROM scfcontrols JOIN scfcontroldetails ON scfcontrols.uuid = scfcontroldetails.control_uuid WHERE scfcontroldetails.control_property = '" + framework + "' AND scfcontrols.scf_domain ='" + domain + "'"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var control models.SCFcontrol
		err = results.Scan(&control.Uuid, &control.Scf_control, &control.Scf_domain, &control.Scf_ref, &control.Control_question, &control.Control_framework, &control.Mapping_values)
		if err != nil {
			panic(err.Error())
		}
		//control.Control_details = Select_control_details_with_filter(control.Uuid, framework)
		controls = append(controls, control)

	}
	defer results.Close()

	return controls

}

func Select_all_controls_per_framework(framework string) []models.SCFcontrol {

	var controls []models.SCFcontrol

	q := "SELECT scfcontrols.uuid, scfcontrols.scf_control, scfcontrols.scf_domain,scfcontrols.scf_ref,scfcontrols.control_question,scfcontroldetails.control_property,scfcontroldetails.control_property_value FROM  scfcontrols JOIN scfcontroldetails ON scfcontrols.uuid = scfcontroldetails.control_uuid WHERE scfcontroldetails.control_property = '" + framework + "'"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)

	defer results.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var control models.SCFcontrol
		err = results.Scan(&control.Uuid, &control.Scf_control, &control.Scf_domain, &control.Scf_ref, &control.Control_question, &control.Control_framework, &control.Mapping_values)
		if err != nil {
			panic(err.Error())
		}
		controls = append(controls, control)

	}

	return controls

}

func Select_department_per_client(companyID string) []models.Department {

	var dpmts []models.Department

	q := "SELECT  name, companyID FROM `departments` WHERE  companyID ='" + companyID + "'"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var dpmt models.Department
		err = results.Scan(&dpmt.Name, &dpmt.CompanyID)
		if err != nil {
			panic(err.Error())
		}

		dpmts = append(dpmts, dpmt)

	}
	defer results.Close()

	return dpmts

}

func Select_users_per_department(companyID string, department string) []models.OctopusUser {

	var users []models.OctopusUser

	q := "SELECT  firstname, lastname, email ,Department, position, user_role,on_leave FROM `users` WHERE  company_id ='" + companyID + "' AND Department = '" + department + "'"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var user models.OctopusUser
		err = results.Scan(&user.Firstname, &user.Lastname, &user.Email, &user.Department, &user.Position, &user.User_role, &user.On_leave)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)

	}
	defer results.Close()

	return users

}

func Insert_evidence_request(evReq *models.EvidenceRequest) {

	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `evidenceRequests` (`req_reference`, `req_owner`, `req_assessor`,`req_reviewer`,`req_status`,`contributors`,`company_id`) VALUES (?, ?, ?, ?, ?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, &evReq.Req_reference, &evReq.Req_owner, &evReq.Req_assessor, &evReq.Req_reviewer, &evReq.Req_status, &evReq.Contributors, &evReq.Company_id)
	fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)
	if err != nil {
		fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Framework: %s", err)
	}

	defer db.Close()

}

func Select_evidence_requests(companyID string) []models.EvidenceRequest {

	var evidreqs []models.EvidenceRequest

	q := "SELECT  req_reference,req_owner,req_assessor,req_reviewer,req_status,contributors FROM `evidencerequests` WHERE  company_id ='" + companyID + "'"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var evidreq models.EvidenceRequest
		var contributors string
		err = results.Scan(&evidreq.Req_reference, &evidreq.Req_owner, &evidreq.Req_assessor, &evidreq.Req_reviewer, &evidreq.Req_status, &contributors)
		evidreq.Contributors = strings.Fields(contributors)
		if err != nil {
			panic(err.Error())
		}

		evidreqs = append(evidreqs, evidreq)

	}
	defer results.Close()

	return evidreqs

}

func Select_evidence_request_controls(reference string) []models.SCFcontrol {

	var controls []models.SCFcontrol

	q := "SELECT   scfcontrols.uuid, scfcontrols.scf_control ,scfcontrols.scf_domain , scfcontrols.scf_ref , scfcontrols.control_question FROM `scfcontrols` LEFT JOIN `deployedcontrols` ON deployedcontrols.control_uuid = scfcontrols.uuid WHERE  deployedcontrols.evidenceReq_ref ='" + reference + "' AND deployedcontrols.control_uuid = scfcontrols.uuid"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var control models.SCFcontrol
		err = results.Scan(&control.Uuid, &control.Scf_control, &control.Scf_domain, &control.Scf_ref, &control.Control_question)
		if err != nil {
			panic(err.Error())
		}
		control.Control_details = Select_control_details(control.Uuid)
		controls = append(controls, control)

	}
	defer results.Close()

	return controls

}

func Insert_deployed_control(reqref string, uuid string) {

	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `deployedcontrols` (`evidenceReq_ref`, `control_uuid`) VALUES ( ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, &reqref, &uuid)
	fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)
	if err != nil {
		fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Framework: %s", err)
	}

	defer db.Close()

}

func Insert_into_library(scope_object *models.ConfirmScope, control_uuid string) {

	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `controllibrary` (`control_uuid`, `request_owner`,`company_id`) VALUES ( ?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, &control_uuid, &scope_object.Req_owner, &scope_object.Company_id)
	fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)
	if err != nil {
		fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Framework: %s", err)
	}

	defer db.Close()

}

func List_controls_from_library(company_id string) []models.SCFcontrol {

	var controls []models.SCFcontrol
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Distinct("scfcontrols.uuid").Select("scfcontrols.uuid", "scfcontrols.scf_control", "scfcontrols.scf_domain", "scfcontrols.scf_ref", "scfcontrols.control_question", "scfcontrols.description").Table("scfcontrols").Joins("JOIN controllibrary ON scfcontrols.uuid = controllibrary.control_uuid").Where("controllibrary.company_id = ?", company_id).Find(&controls)

	var frameworks []string
	for i := 0; i < len(controls); i++ {
		fmt.Println(i)
		db.Distinct("scfcontroldetails.control_property").Select("scfcontroldetails.control_property").Table("scfcontroldetails").Where("scfcontroldetails.control_uuid = ? ", controls[i].Uuid).Find(&frameworks)
		controls[i].Control_details = frameworks
		frameworks = nil
	}
	return controls

}

func Select_control_join_details(search_word string) []models.SCFcontrol {

	var controls []models.SCFcontrol
	q := "SELECT DISTINCT scfcontrols.uuid, scfcontrols.scf_control, scfcontrols.scf_domain, scfcontrols.scf_ref, scfcontrols.control_question FROM scfcontrols INNER JOIN scfcontroldetails ON scfcontrols.uuid = scfcontroldetails.control_uuid WHERE scfcontroldetails.control_property LIKE '%" + search_word + "%'"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	results, err := db.Query(q)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {

		var control models.SCFcontrol
		err = results.Scan(&control.Uuid, &control.Scf_control, &control.Scf_domain, &control.Scf_ref, &control.Control_question)
		if err != nil {
			panic(err.Error())
		}
		control.Control_details = Select_control_details_with_filter(control.Uuid, search_word)
		controls = append(controls, control)

	}
	defer results.Close()

	return controls

}

func Select_control_count(reference string) int {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	// Query for a value based on a single row.
	results, err := db.Query("SELECT COUNT(*) FROM scfcontrols JOIN scfcontroldetails ON scfcontrols.uuid = scfcontroldetails.control_uuid WHERE scfcontroldetails.control_property='" + reference + "'")
	var control_count int
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {
		err = results.Scan(&control_count)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("================ Printing controls_count on 711 =================")
		fmt.Println(control_count)
	}
	return control_count
}

func Select_domains_count(reference string) int {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	// Query for a value based on a single row.
	results, err := db.Query("SELECT COUNT(DISTINCT scf_domain) FROM scfcontrols JOIN scfcontroldetails ON scfcontrols.uuid = scfcontroldetails.control_uuid WHERE scfcontroldetails.control_property='" + reference + "'")
	var domain_count int
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {
		err = results.Scan(&domain_count)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("================ Printing controls_count on 768 =================")
		fmt.Println(domain_count)
	}
	return domain_count
}

func Select_controls_from_selected_frameworks(frms []string, domain string) []models.SCFcontrol {

	var controls []models.SCFcontrol
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Distinct("scfcontrols.uuid").Select("scfcontrols.uuid", "scfcontrols.scf_control", "scfcontrols.scf_domain", "scfcontrols.scf_ref", "scfcontrols.control_question").Table("scfcontrols").Joins("JOIN scfcontroldetails ON scfcontrols.uuid = scfcontroldetails.control_uuid").Where("scfcontroldetails.control_property IN ? AND scfcontrols.scf_domain= ?", frms, domain).Find(&controls)
	fmt.Println("====== printing controls array size ============")
	fmt.Println(len(controls))

	fmt.Println("====== done printing !============")
	var frameworks []string
	for i := 0; i < len(controls); i++ {
		fmt.Println(i)
		db.Distinct("scfcontroldetails.control_property").Select("scfcontroldetails.control_property").Table("scfcontroldetails").Where("scfcontroldetails.control_property IN ?  AND  scfcontroldetails.control_uuid = ? ", frms, controls[i].Uuid).Find(&frameworks)
		controls[i].Control_details = frameworks
		frameworks = nil
	}

	fmt.Println("Printing lengh for " + domain)
	fmt.Println(len(controls))
	return controls

}
