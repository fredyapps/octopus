package models

import "gorm.io/gorm"

type Fields struct {
	gorm.Model
	Fields []Field `json:"fields"`
}
