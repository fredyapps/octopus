//  "UUID": "857e0d76-e6b2-4893-b901-e8bf6518a8c8",
//  "#\u00a0": 1,
//  "SCF Domain": "Security & Privacy Governance",
//  "SCF Identifier": "GOV",
//  "Security & Privacy by Design (S|P) Principles":
//  Principle Intent

package models

import (
	"gorm.io/gorm"
)

type SCFDomain struct {
	gorm.Model
	Id_domain       int      `json:"id_domain"`
	UUID            string   `json:"UUID"`
	ID              int      `json:"#\u00a0"`
	SCFDomain       string   `json:"SCF Domain"`
	SCFIdentifier   string   `json:"SCF Identifier"`
	SecurityPrivacy []string `json:"Security & Privacy by Design (S|P) Principles"`
	PrincipleIntent string   `json:"Principle Intent"`
	Date_created    string   `gorm:"-:all"`
	Date_updated    string   `json:"Date_updated"`
}