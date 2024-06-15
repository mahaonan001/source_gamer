package router

import (
	"source_gamer/controller"
	"source_gamer/middle"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {

	api := r.Group("/api")
	{
		player := api.Group("/user")
		{
			player.POST("/register", controller.Register)
			player.POST("/login", controller.Login)
			player.POST("/info", middle.AuthMiddleware(), controller.Info)
			player.POST("/cg_info", middle.AuthMiddleware(), controller.Cgif)
			// player.POST("/upload")
			player.POST("/g_record", middle.AuthMiddleware(), controller.Get_records)
			player.GET("/location", controller.Locations)
			chat := player.Group("/chat")
			{
				chat.POST("/", middle.AuthMiddleware(), controller.Chats)
				chat.POST("/records", middle.AuthMiddleware(), controller.ChatsRecord)
			}

		}
		api.POST("/get_code", controller.Send_Code)
	}
	return r
}
