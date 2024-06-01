package router

import (
	"source_gamer/common"
	"source_gamer/controller"
	"source_gamer/middle"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	common.Init_db()
	api := r.Group("/api")
	{
		player := api.Group("/user")
		{
			player.POST("/register", controller.Register)
			player.POST("/login", controller.Login)
			player.GET("/info", middle.AuthMiddleware(), controller.Info)
			player.POST("/cg_info", middle.AuthMiddleware(), controller.Cgif)
			player.POST("/upload")
		}
		api.POST("/get_code", controller.Send_Code)
	}
	return r
}
