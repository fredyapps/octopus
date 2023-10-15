package models

//"gorm.io/gorm"

type Framework struct {
	Id_framework int    `json:"id"`
	Name         string `json:"name"`
	Reference    string `json:"reference"`
	Version      string `json:"version"`
	Description  string `json:"description"`
	Created_by   string `json:"created_by"`
	Created_at   string `json:"created_at"`
}
