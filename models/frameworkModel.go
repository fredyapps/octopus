package models

//"gorm.io/gorm"

type Framework struct {
	Id_framework  int    `json:"id"`
	Name          string `json:"name"`
	Reference     string `json:"reference"`
	Version       string `json:"version"`
	Numb_controls int    `json:"numb_controls"`
	Numb_layers   int    `json:"numb_layers"`
	Description   string `json:"description"`
	Created_by    string `json:"created_by"`
	Created_at    string `json:"created_at"`
}
