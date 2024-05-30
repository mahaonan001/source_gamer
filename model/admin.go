package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email      string `gorm:"type:varchar(20);not null;primaryKey"`
	Name       string `gorm:"type:varchar(20);not null"`
	PassWord   string `gorm:"type:varchar(255);not null"`
	ErrorTimes int8   `gorm:"type:int8;not null"`
}
