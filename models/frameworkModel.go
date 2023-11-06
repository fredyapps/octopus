package models

//"gorm.io/gorm"

type Framework struct {
	Id_framework       int    `json:"id"`
	Name               string `json:"name"`
	Reference          string `json:"reference"`
	Version            string `json:"version"`
	Description        string `json:"description"`
	Number_of_controls int
	Number_of_domains  int
}
