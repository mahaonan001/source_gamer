package controller

import (
	"log"
	"net/http"
	"source_gamer/common"
	"source_gamer/model"
	"source_gamer/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	//获取参数
	Email := c.PostForm("email")
	PassWord := c.PostForm("password")
	//数据验证
	if len(PassWord) < 6 {
		response.Response(c, http.StatusOK, 400, nil, "密码不能低于6位数")
		return
	}

	db, _ := common.GetDB()

	log.Println("email:", Email, "PassWord:", PassWord, " is logining")
	User := isEmailExited(db, Email)
	log.Println(User.ID)
	if User.ID == 0 {
		response.FalseRe(c, "用户不存在", nil)
		return
	} else if User.ErrorTimes >= 3 {
		response.FalseRe(c, "密码错误次数过多，账号已冻结", nil)
		return
	}
	//验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(User.PassWord), []byte(PassWord)); err != nil {
		response.FalseRe(c, "密码错误", nil)
		User.ErrorTimes++
		db.Save(&User)
		return
	}
	token, err := common.ReleaseToken_User(User)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "token生成失败，系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	DB_code, _ := common.GetDB()
	DB_code.Where("email=?", User.Email).Delete(&model.EmailCode{})
	response.SuccessRe(c, "登陆成功", gin.H{"token": token})
	User.ErrorTimes = 0
	db.Save(&User)
}
func isEmailExited(db *gorm.DB, Email string) model.User {
	var User model.User
	db.Where("Email = ?", Email).Find(&User)
	return User
}
