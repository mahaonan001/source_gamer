package controller

import (
	"source_gamer/common"
	"source_gamer/response"

	"github.com/gin-gonic/gin"
)

type Location struct {
	Ip  string  `gorm:"column:ip"`
	Pos float32 `gorm:"column:pos"`
	Neg float32 `gorm:"column:neg"`
}

func (l Location) TableName() string {
	return "locations"
}

func Locations(c *gin.Context) {
	var locations []Location
	db, _ := common.GetDB()
	db.Order("pos desc").Find(&locations)
	response.SuccessRe(c, "", gin.H{"locations": locations})
}
