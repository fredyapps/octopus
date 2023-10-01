package models

type OctopusUser struct {
	//gorm.Model

	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	Department   string `json:"department"`
	Position     string `json:"position"`
	User_role    string `json:"user_role"`
	On_leave     string `json:"on_leave"`
	Staff_number string `json:"staff_number"`
	Phone_number string `json:"phone_number"`
	Company_id   string `json:"company_id"`
	Date_created string `json:"date_created"`
}
