package model

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	V_type           string
	Coding           string
	V_link           string
	Page_n           int64
	User_name        string
	User_id          string
	User_home        string
	Time             string
	Ip               string
	Like_n           int64
	Like_l           string
	Cleaned_comments string
}
