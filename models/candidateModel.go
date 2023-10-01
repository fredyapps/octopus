package models

import (
	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	Id_contact int
	FullName   string `json:"fullName"`
	Subject    string `json:"subject"`
	Email      string `json:"email"`
	Message    string `json:"Message"`
	Created_at string
}
