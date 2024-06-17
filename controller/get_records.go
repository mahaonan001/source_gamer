package controller

import (
	"fmt"
	"source_gamer/common"
	"source_gamer/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Show struct {
	ID               string  `gorm:"column:record_id"`
	Cleaned_comments string  `gorm:"column:cleaned_comments"`
	Option_word      string  `gorm:"column:option_word"`
	Score_           bool    `gorm:"column:score_"`
	T_room           float64 `gorm:"column:t_room"`
	S_room           int     `gorm:"column:s_room"`
	HotDevice        string  `gorm:"column:hotdevice"`
	BurnningT        string  `gorm:"column:burnning_t"`
	Device_logo      string  `gorm:"column:device_logo"`
	Hot_T            string  `gorm:"column:hot__t"`
	Time_cyc         string  `gorm:"column:time_cyc"`
	Money_cyc        float64 `gorm:"column:money_cyc"`
	Gas_cyc          float64 `gorm:"column:gas_cyc"`
	Ele_cyc          int     `gorm:"column:ele_cyc"`
	Boal_cyc         int     `gorm:"column:boal_cyc"`
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
	db.Offset(begin).Limit(10).Find(&records, "chat=?", 0)
	if db.Error != nil {
		response.FalseRe(c, "数据库出错", nil)
	}
	response.SuccessRe(c, "", gin.H{"data": records})
}
func SearchRecord(c *gin.Context) {
	comment := c.PostForm("comment")
	db, err := common.GetDB()
	if err != nil {
		return
	}
	var shows []Show
	comment_ := fmt.Sprintf("%%%s%%", comment)
	result := db.Where("cleaned_comments like ?", comment_).Find(&shows)
	if result.Error != nil {
		response.FalseRe(c, "some err happend", gin.H{"err": result.Error})
	}
	response.SuccessRe(c, "success", gin.H{"data": shows})
}
