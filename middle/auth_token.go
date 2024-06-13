package middle

import (
	"net/http"
	"source_gamer/common"
	"source_gamer/model"
	"source_gamer/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{"code": 401, "msg": "权限不足"}, "")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{"code": 401, "msg": "权限不足 "}, "")
			ctx.Abort()
			return
		}

		//验证通过，获取UserID
		UserId := claims.UserId
		DB := common.GetDB()
		var User model.User
		DB.First(&User, UserId)

		//用户信息不存在
		if User.ID == 0 {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{"code": 401, "msg": "权限不足 "}, "")
			ctx.Abort()
			return
		}
		//用户存在，将User信息写入上下文
		ctx.Set("User", User)
		ctx.Next()
	}
}
