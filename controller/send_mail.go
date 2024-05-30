package controller

import (
	"source_gamer/mail"

	"github.com/gin-gonic/gin"
)

func Send_Code(c *gin.Context) {
	arm_email := c.PostForm("email")
	mail.SendMail(arm_email, c)
}
