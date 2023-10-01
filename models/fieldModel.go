package models

type Field struct {
	//gorm.Model
	Name    string `json:"name"`
	Type    string `json:"type"`
	Segment string `json:"segment"`
}
