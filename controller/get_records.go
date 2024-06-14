package controller

import (
	"source_gamer/common"
	"source_gamer/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Show struct {
	ID               string
	Cleaned_comments string
	Option_word      string
	Score_           bool
	T_room           float64
	S_room           int
	BurnningT        string
	Device_logo      string
	Hot_T            string
	Time_cyc         string
	Money_cyc        float64
	Gas_cyc          float64
	Ele_cyc          int
	Boal_cyc         int
}

func (s Show) TableName() string {
	return "shows"
}
func Get_records(c *gin.Context) {
	begin_ := c.PostForm("begin")
	db, _ := common.GetDB()
	var records []Show
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
