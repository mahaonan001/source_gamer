package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"type:varchar(20);not null;primary Key"`
	Name       string `gorm:"type:varchar(20);not null"`
	PassWord   string `gorm:"type:varchar(255);not null"`
	ErrorTimes int8   `gorm:"type:int8;not null"`
}

type EmailCode struct {
	ID         uint64
	Email      string    `gorm:"type:varchar(30);not null;primary Key"`
	Code_email string    `gorm:"type:varchar(6);not null;unique"`
	InfTime    time.Time `gorm:"type:datetime;not null"`
}