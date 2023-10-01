package models

import "gorm.io/gorm"

type Schema struct {
	gorm.Model
	Id_field int
	Segment  string
	Name     string
	Type     string
}
