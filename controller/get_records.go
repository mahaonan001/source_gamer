package controller

import (
	"source_gamer/common"
	"source_gamer/model"
	"source_gamer/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get_records(c *gin.Context) {
	begin_ := c.PostForm("begin")
	db := common.GetDB()
	var records []model.Show
	begin, err := strconv.Atoi(begin_)
	if err != nil {
		response.FalseRe(c, "页数错误", nil)
	}
	db.Offset(begin).Limit(10).Find(&records)
	if db.Error != nil {
		response.FalseRe(c, "数据库出错", nil)
	}
	response.SuccessRe(c, "", gin.H{"data": records})
}
