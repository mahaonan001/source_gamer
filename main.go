package main

import (
	"log"
	"os"
	"source_gamer/common"
	"source_gamer/router"
	"source_gamer/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	tick := time.NewTicker(time.Duration(viper.GetInt("TimeCyc")) * time.Second)
	go func() {
		for range tick.C {
			utils.Timely()
		}
	}()
	r.Use(cors.New(cors.Config{
		// 允许的域名或IP地址
		AllowOrigins: []string{"*"},
		// 允许的请求方法（GET, POST等）
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// 允许的请求头
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		// 是否允许认证信息跟随请求
		AllowCredentials: true,
	}))
	r = router.CollectRouter(r)
	r.Run(":" + viper.GetString("server.port"))
}
func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	db, err := common.GetDB()
	if err != nil {
		err = common.Init_db()
		if err != nil {
			log.Panicln(err)
		}
		sqlLocatuon, err := os.ReadFile(viper.GetString("SQL.locations"))
		if err != nil {
			panic("failed to read sql file")
		}
		sqlShow, err := os.ReadFile(viper.GetString("SQL.shows"))
		if err != nil {
			panic("failed to read sql file")
		}
		db.Exec(string(sqlLocatuon))
		db.Exec(string(sqlShow))
		utils.Timely()
	}

}
