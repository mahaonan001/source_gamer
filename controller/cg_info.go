package controller

import (
	"source_gamer/common"
	"source_gamer/model"
	"source_gamer/response"

	"github.com/gin-gonic/gin"
)

func Cgif(c *gin.Context) {
	getA, _ := c.Get("Admin")
	var User model.User
	if getA != nil {
		email := c.PostForm("email")
		User.Email = email
	}
	getU, _ := c.Get("User")
	if getU != nil {
		User = getU.(model.User)
	}
	name := c.PostForm("name")
	db, _ := common.GetDB()
	User.Name = name
	db.Save(User)
	if db.Error != nil {
		response.FalseRe(c, "err to change info", gin.H{"before": getU, "after": User})
		return
	}
	response.SuccessRe(c, "update info successfully", nil)
}
