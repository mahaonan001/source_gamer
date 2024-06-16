package controller

import (
	"log"
	"net/http"
	"regexp"
	"source_gamer/common"
	"source_gamer/model"
	"source_gamer/response"

	"source_gamer/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register 用户注册函数
func Register(c *gin.Context) {
	//获取参数
	db_User, _ := common.GetDB()
	Email := c.PostForm("email")
	Name := c.PostForm("name")
	PassWord := c.PostForm("password")
	Code_Email := c.PostForm("code")
	//数据验证

	if len(PassWord) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 403, nil, "密码不能低于6位数")
		return
	}
	if len(Name) == 0 {
		Name = utils.RandomString(10, "qwertyuiopasdfghjklzxcvbnm1234567890")
	}

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	// 使用MatchString()函数来判断电子邮件地址是否匹配正则表达式
	if emailRegex.MatchString(Email) {
		log.Println("name:", Name, "email:", Email)
		if isEmailExited_User(db_User, Email).ID != 0 {
			response.Response(c, http.StatusOK, 403, nil, "该邮箱已注册")
			return
		}
		DB_code, _ := common.GetDB()
		var Code_email model.EmailCode
		DB_code.Where("email=?", Email).Order("id desc").Limit(1).Find(&Code_email)
		if Code_Email == Code_email.Code_email { //存在符合条件的验证码
			HashPassword, err := bcrypt.GenerateFromPassword([]byte(PassWord), bcrypt.DefaultCost)
			if err != nil {
				response.Response(c, http.StatusInternalServerError, 500, nil, "加密失败")
				log.Println(err)
				return
			}
			newUser := model.User{
				Name:       Name,
				PassWord:   string(HashPassword),
				ErrorTimes: 0,
				Email:      Email,
			}
			db_User.Create(&newUser)
			response.SuccessRe(c, "注册成功", gin.H{"code": 200, "msg": "注册成功"})
		} else {
			response.Response(c, http.StatusOK, 402, gin.H{"msg": "验证码填写错误"}, "err")
		}
	}
}

// isEmailExited_User 检查邮箱是否存在对应的用户
func isEmailExited_User(db *gorm.DB, Email string) model.User {
	var User model.User
	db.Where("Email = ?", Email).Order("id asc").Limit(1).Find(&User)
	return User
}
