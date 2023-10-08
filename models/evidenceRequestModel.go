package models

import (
	"gorm.io/gorm"
)

type EvidenceRequest struct {
	gorm.Model
	Id_request    int         `json:"id_request"`
	Req_reference string      `json:"req_reference"`
	Req_owner     string      `json:"req_owner"`
	Req_assessor  string      `json:"req_assessor"`
	Req_reviewer  string      `json:"req_reviewer"`
	Req_status    string      `json:"req_status"`
	Req_progress  string      `json:"req_progress"`
	Contributors  interface{} `json:"contributors"`
	Controls      interface{} `json:"controls"`
	Company_id    string      `json:"company_id"`
	Date_created  string      `json:"date_created"`
	Date_updated  string      `json:"date_updated"`
}

// {
//     "req_owner": "the owner",
//     "req_assessor": " the assessor ",
//     "req_reviewer": " the reviewer",
//     "req_status": " pending",
//     "contributors": "fredysallah@gmail.com,solutions@gmail.com, yaa.koateng@hotmail.com",
//     "controls": {},
//     "company_id": "SSNIT_GHANA"
// }
