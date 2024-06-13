package controller

import (
	"source_gamer/common"
	"source_gamer/response"

	"github.com/gin-gonic/gin"
)

type Location struct {
	zone string
	po   float32
	ne   float32
}

func Locations(c *gin.Context) {
	var locations Location
	db := common.GetDB()
	db.Order("pos").Find(&locations)
	response.SuccessRe(c, "", gin.H{"locations": locations})
}
