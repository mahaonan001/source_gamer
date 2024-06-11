package model

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	V_type           string `gorm:"type:varchar(10);"`
	Coding           string `gorm:"type:varchar(10);"`
	V_link           string `gorm:"type:varchar(100);"`
	Page_n           int64  `gorm:"type:bigint;"`
	User_name        string `gorm:"type:varchar(20)"`
	User_id          string `gorm:"type:varchar(20)"`
	User_home        string `gorm:"type:varchar(100)"`
	Time             string `gorm:"type:datetime"`
	Ip               string `gorm:"type:varchar(10)"`
	Like_n           int64  `gorm:"type:bigint"`
	Like_l           string `gorm:"type:varchar(5)"`
	Cleaned_comments string `gorm:"type:longtext;not null"`
}

type Score struct {
}

type Keyword struct {
}

type Dim struct {
}
