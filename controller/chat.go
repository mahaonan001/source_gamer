package controller

import (
	"log"
	"source_gamer/common"
	"source_gamer/model"
	"source_gamer/response"
	"source_gamer/utils"

	"github.com/gin-gonic/gin"
)

func Chats(c *gin.Context) {
	User, t := c.Get("User")
	if !t {
		response.FalseRe(c, "身份信息有误，请重新登陆", nil)
	}
	db, err := common.GetDB()
	if err != nil {
		response.FalseRe(c, "数据库连接失败", nil)
	}
	comment := c.PostForm("comment")
	var randId string
	for {
		randId = utils.RandomString(9, "1234567890qwertyuiopasdfghjklzxcvbnm")
		result := db.Select("id=?", randId)
		if result.RowsAffected == 0 {
			break
		}
	}

	record := model.Record{
		ID:               randId,
		Chat:             true,
		Cleaned_comments: comment,
	}
	db.Create(&record)
	log.Println(User)
	chat := model.Chat{
		EmailId:  User.(model.User).Email,
		RecordId: randId,
	}
	db.Create(&chat)
	response.SuccessRe(c, "聊天记录", gin.H{"id": randId})
}

func ChatsRecord(c *gin.Context) {
	User, t := c.Get("User")
	if !t {
		response.FalseRe(c, "身份信息有误，请重新登陆", nil)
	}
	db, err := common.GetDB()
	if err != nil {
		response.FalseRe(c, "数据库连接失败", nil)
	}
	var chats []model.Chat
	db.Where("email_id=?", User.(model.User).Email).Find(&chats)

	var records []string
	for _, chat := range chats {
		var record model.Record
		db.Where("id=? and chat=?", chat.RecordId, true).Find(&record)
		records = append(records, record.Cleaned_comments)

	}
	response.SuccessRe(c, "历史记录", gin.H{"历史": records})
}
