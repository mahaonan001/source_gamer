package controller

import (
	"log"
	"net/http"
	"source_gamer/common"
	"source_gamer/model"
	"source_gamer/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func A_login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if len(password) < 6 {
		response.Response(c, http.StatusOK, 400, nil, "密码不能低于6位数")
		return
	}
	db := common.GetDB_Admin()
	var admin model.Admin
	db.Where("email = ?", account).Find(&admin)
	if admin.ID == 0 {
		response.FalseRe(c, "登录错误", gin.H{"information": "账号或密码错误"})
		return
	} else if admin.ErrorTimes >= 3 {
		response.FalseRe(c, "密码错误次数过多，账号已冻结", nil)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.PassWord), []byte(password)); err != nil {
		response.FalseRe(c, "密码错误", nil)
		admin.ErrorTimes++
		db.Save(&admin)
		return
	}
	token, err := common.ReleaseToken(admin)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "token生成失败，系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	DB_code := common.GetDB_Email()
	DB_code.Where("email=?", admin.Email).Delete(&model.EmailCode{})
	response.SuccessRe(c, "登陆成功", gin.H{"token": token})
	admin.ErrorTimes = 0
	db.Save(&admin)
}

func Get_Info(c *gin.Context) {
	Admin, _ := c.Get("User")
	db := common.GetDB_User()
	var Users_ []model.User
	db.Find(&Users_)
	response.SuccessRe(c, "successfully get users", gin.H{"admin": Admin, "users": Users_})
}
