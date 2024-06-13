package model

import (
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	V_type           string `gorm:"type:varchar(10);"`
	Coding           string `gorm:"type:varchar(10);"`
	V_link           string `gorm:"type:varchar(100);"`
	Page_n           int64  `gorm:"type:bigint;"`
	User_name        string `gorm:"type:varchar(20)"`
	User_id          string `gorm:"type:varchar(20)"`
	User_home        string `gorm:"type:varchar(255)"`
	Time             string `gorm:"type:varchar(40)"`
	Ip               string `gorm:"type:varchar(10)"`
	Like_n           int64  `gorm:"type:bigint"`
	Like_l           string `gorm:"type:varchar(5)"`
	Cleaned_comments string `gorm:"type:longtext;not null"`
}

type Score struct {
	RecordId        int    `gorm:"index"`
	Record          Record `gorm:"foreignKey:RecordId"`
	Analysis        string `gorm:"type:longtext;not null"`
	Extracted_texts string `gorm:"type:longtext;not null"`
	Dim_id          int    `gorm:"index"`
	Dim             Dim    `gorm:"foreignKey:Dim_id"`
	Option_word     string `gorm:"type:varchar(20);not null"`
	Score_          bool   `gorm:"type:int;not null"`
}

type Keyword struct {
	RecordId int     `gorm:"index"`
	Record   Record  `gorm:"foreignKey:RecordId"`
	T_room   float64 `gorm:"type:double;"`
	S_room   int     `gorm:"type:int"`
	// BurnningT string  `gorm:""`
}

type Dim struct {
	Id   int    `gorm:"type:int;autoincrement;primaryKey;not null"`
	Dim_ string `gorm:"type:varchar(10);not null;unique"`
}

type Chat struct {
	gorm.Model
	Email_   string `gorm:"type:index"`
	User     User   `gorm:"foreignKey:Email_;references:Email"` // Use Email as the reference
	RecordId int    `gorm:"index"`
	Record   Record `gorm:"foreignKey:RecordId;references:ID"`
}
