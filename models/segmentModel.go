package models

// segments := [8]string{
// 	"Domains & Principles",
// 	"SCF 2023.2",
// 	"Assessment Objectives 2023.2",
// 	"Evidence Request List 2023.2",
// 	"Privacy Management 2023.2",
// 	"Risk Catalog",
// 	"Threat Catalog",
// 	"Authoritative Sources"}

type Segment struct {
	//gorm.Model
	Domains_and_principles      Fields `json:"Domains & Principles"`
	SCF20232                    Fields `json:"SCF 2023.2"`
	Assessment_Objectives_20232 Fields `json:"Assessment Objectives 2023.2"`
	Evidence_Request_List_20232 Fields `json:"Evidence Request List 2023.2"`
	Privacy_Management_20232    Fields `json:"Privacy Management 2023.2"`
	Risk_Catalog                Fields `json:"Risk Catalog"`
	Threat_Catalog              Fields `json:"Threat Catalog"`
	Authoritative_Sources       Fields `json:"Authoritative Sources"`
}
