package models

type SCFcontrol struct {
	//gorm.Model
	//Id_scf_control   int
	Uuid              string
	Scf_control       string
	Scf_domain        string
	Scf_ref           string
	Control_question  string
	Control_framework string
	Mapping_values    string
	Control_details   interface{}
}
