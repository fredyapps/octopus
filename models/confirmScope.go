package models

type ConfirmScope struct {
	//gorm.Model

	Req_owner  string      `json:"req_owner"`
	Controls   interface{} `json:"controls"`
	Company_id string      `json:"company_id"`
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
