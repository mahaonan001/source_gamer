package controller

import (
	"source_gamer/model"
	"source_gamer/response"

	"github.com/gin-gonic/gin"
)

type User_Info struct {
	Email string
	Name  string
}

func Info(c *gin.Context) {
	User, _ := c.Get("User")
	response.Response(c, 200, 200, gin.H{"user": User_Info{
		Name:  User.(model.User).Name,
		Email: User.(model.User).Email,
	}}, "successfully get info")
}
