package models

type SCFcontrolDetail struct {
	//gorm.Model
	Id_control_detail      int
	Control_uuid           string
	Control_property       string
	Control_property_value string
}
