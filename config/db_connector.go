package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"octopus/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
)

// Connect to the database
var Database *gorm.DB
var err error

//var dsn string = "root:@tcp(127.0.0.1:3306)/octopus"

var dsn string = "root:Mathapelo@2030@tcp(127.0.0.1:3306)/octopus"

// var db *sql.DB
var insert sql.Result

func Connect() {

	// Open the database.
	db, err := sql.Open("mysql", dsn)

	fmt.Println("================ Error  connector line 35 =================")
	fmt.Println(db)
	fmt.Println(db.Ping())

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

func Select_frameworks() []models.Framework {

	var frameworks []models.Framework

	q := "SELECT  id_framework, description ,reference , name , version FROM `Frameworks`"
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
		fmt.Println(framework)
		frameworks = append(frameworks, framework)

	}
	defer results.Close()

	return frameworks

}

func Insert_framework(frmk *models.Framework) {

	//=========================================== |||||||||||||| ============================================
	//q := "INSERT INTO `Frameworks` (`name`, `reference`, `version`,`numb_controls`,`numb_layers`,`description`) VALUES ('Ananisallah', 'insertion', 'fredy@gmail.com', 'holla')"
	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `Frameworks` (`name`, `reference`, `version`,`numb_controls`,`numb_layers`,`description`) VALUES (?, ?, ?, ?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, frmk.Name, frmk.Reference, frmk.Version, frmk.Numb_controls, frmk.Numb_layers, frmk.Description)
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

func Insert_domain(dmn *models.SCFDomain) {

	//=========================================== |||||||||||||| ============================================
	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `SCFDomains` (`UUID`, `ID`, `SCFDomain`,`SCFIdentifier`,`SecurityPrivacy`,`PrincipleIntent`) VALUES (?, ?, ?, ?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, dmn.UUID, dmn.ID, dmn.SCFDomain, dmn.SCFIdentifier, dmn.SecurityPrivacy[0], dmn.PrincipleIntent)
	fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)

	if err != nil {
		fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Domain: %s", err)
	}

	defer db.Close()

}

func Insert_schema(fld models.Field) {

	//=========================================== |||||||||||||| ============================================
	db, err := sql.Open("mysql", dsn)
	query := "INSERT INTO `JSON_schema_fields` (`segment`, `name`, `type`) VALUES (?, ?, ?)"

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
	query := "INSERT INTO `SCFcontrols` (`uuid`, `scf_control`, `scf_domain`,`scf_ref`,`control_question`) VALUES (?, ?, ?, ?, ?)"

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
	query := "INSERT INTO `SCFcontrolDetails` (`control_uuid`, `control_property`, `control_property_value`) VALUES (?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, ctrlDet.Control_uuid, ctrlDet.Control_property, ctrlDet.Control_property_value)
	//fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)

	if err != nil {
		//fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Domain: %s", err)
	}

	defer db.Close()

}

func Select_tariff_plans() []models.TariffPlan {

	var plans []models.TariffPlan

	q := "SELECT  plan, reference  FROM `tariff_plans`"
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

		var plan models.TariffPlan
		err = results.Scan(&plan.Plan, &plan.Reference)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(plan)
		plans = append(plans, plan)

	}
	defer results.Close()

	return plans

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

	q := "SELECT  uuid, scf_control ,scf_domain , scf_ref , control_question FROM `SCFcontrols`"

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

	q := "SELECT  uuid, scf_control ,scf_domain , scf_ref , control_question FROM `SCFcontrols` WHERE  scf_domain ='" + domain + "'"
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

	q := "SELECT  firstname, lastname ,email , department , position , user_role , on_leave  FROM `Users` WHERE  company_id ='" + comp_id + "'"
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

//SELECT  uuid, scf_control ,scf_domain , scf_ref , control_question,(SELECT control_property , control_property_value  FROM `SCFcontrolDetails` WHERE  control_uuid = SCFcontrols.uuid) as control_details FROM `SCFcontrols`

func Select_control_details(uuid string) []models.SCFcontrolDetail {

	var ctrl_details []models.SCFcontrolDetail
	q := "SELECT control_property , control_property_value  FROM `SCFcontrolDetails` WHERE  control_uuid ='" + uuid + "'"
	db, err := sql.Open("mysql", dsn)

	results, err := db.Query(q)

	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	for results.Next() {

		var ctrl_detail models.SCFcontrolDetail
		err = results.Scan(&ctrl_detail.Control_property, &ctrl_detail.Control_property_value)

		ctrl_details = append(ctrl_details, ctrl_detail)
	}

	defer results.Close()

	return ctrl_details

}

func Select_controls_with_details_per_domain(domain string) []models.SCFcontrol {

	var controls []models.SCFcontrol

	q := "SELECT  uuid, scf_control ,scf_domain , scf_ref , control_question FROM `SCFcontrols` WHERE  scf_domain ='" + domain + "'"
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

func Select_all_controls_with_details() []models.SCFcontrol {

	var controls []models.SCFcontrol

	q := "SELECT  uuid, scf_control ,scf_domain , scf_ref , control_question FROM `SCFcontrols`"
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

	q := "SELECT  firstname, lastname, email ,Department, position, user_role,on_leave FROM `Users` WHERE  company_id ='" + companyID + "' AND Department = '" + department + "'"
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

	q := "SELECT  req_reference,req_owner,req_assessor,req_reviewer,req_status,contributors FROM `evidenceRequests` WHERE  company_id ='" + companyID + "'"
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

	q := "SELECT   SCFcontrols.uuid, SCFcontrols.scf_control ,SCFcontrols.scf_domain , SCFcontrols.scf_ref , SCFcontrols.control_question FROM `SCFcontrols` LEFT JOIN `deployedControls` ON deployedControls.control_uuid = SCFcontrols.uuid WHERE  deployedControls.evidenceReq_ref ='" + reference + "' AND deployedControls.control_uuid = SCFcontrols.uuid"
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
	query := "INSERT INTO `deployedControls` (`evidenceReq_ref`, `control_uuid`) VALUES ( ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), query, &reqref, &uuid)
	fmt.Println("================ Error  connector line 88 =================")
	fmt.Println(insertResult)
	if err != nil {
		fmt.Println("================ Error  connector line 91 =================")
		log.Fatalf("impossible insert Framework: %s", err)
	}

	defer db.Close()

}

// SELECT SCFcontrols.uuid, SCFcontrols.scf_control, SCFcontrols.scf_domain, SCFcontrolDetails.control_property,SCFcontrolDetails.control_property_value
// FROM SCFcontrols
// LEFT JOIN SCFcontrolDetails
// ON SCFcontrols.uuid = SCFcontrolDetails.control_uuid;
