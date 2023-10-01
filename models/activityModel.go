package models

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Id_activity      int    `json:"id_activity"`
	Framework_id     int    `json:"framework_id"`
	Activity_ower    string `json:"activity_ower"`
	Date_start       string `json:"date_start"`
	Date_end         string `json:"date_end"`
	Review_frequency string `json:"review_frequency"`
	First_approval   string `json:"first_approval"`
	Second_approval  string `json:"second_approval"`
	Third_approval   string `json:"third_approval"`
	Response_time    string `json:"response_time"`
	Activity_status  string `json:"activity_status"`
	Client_company   string `json:"client_company"`
	Progress_status  int    `json:"progress_status"`
	Date_created     string `json:"date_created"`
}
