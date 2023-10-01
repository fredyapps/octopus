package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model

	Company_name   string `json:"company_name"`
	Country        string `json:"country"`
	Contact_person string `json:"contact_person"`
	Contact_email  string `json:"contact_email"`
	Tariff_plan    string `json:"tariff_plan"`
}
