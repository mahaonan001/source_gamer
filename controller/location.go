package controller

import (
	"source_gamer/common"
	"source_gamer/model"
	"source_gamer/response"

	"github.com/gin-gonic/gin"
)

func Locations(c *gin.Context) {
	var locations []model.Location
	db, _ := common.GetDB()
	db.Order("pos desc").Find(&locations)
	response.SuccessRe(c, "", gin.H{"locations": locations})
}
