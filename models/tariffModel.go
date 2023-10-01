package models

type TariffPlan struct {
	//gorm.Model
	//Id_tariff    string `json:"name"`
	Plan      string `json:"plan"`
	Reference string `json:"reference"`
}
