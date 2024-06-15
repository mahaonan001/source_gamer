package main

import (
	"fmt"
	"log"
	"os"
	"source_gamer/common"
	"source_gamer/model"
	"source_gamer/router"
	"source_gamer/utils"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	open := viper.GetBool("TimeSwitch")
	if open {
		db, _ := common.GetDB()
		tick := time.NewTicker(time.Duration(viper.GetInt("TimeCyc")) * time.Second)
		go func() {
			for range tick.C {
				utils.Timely(db)
			}
		}()
	}
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
	err := r.Run(":" + viper.GetString("server.port"))
	if err != nil {
		return
	}
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
	viper2 := viper.New()
	viper2.SetConfigName("first")
	viper2.SetConfigType("yml")
	viper2.AddConfigPath(workDir + "/config")
	err2 := viper2.ReadInConfig()
	if err2 != nil {
		panic(err2)
	}
	var config model.Config
	err2 = viper2.UnmarshalExact(&config)
	if err2 != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err2))
	}
	if config.FirstTime {
		errInit := common.Init_db()
		if errInit != nil {
			log.Panicln("err", errInit)
			return
		}
		sqlLocation, err := os.ReadFile(viper.GetString("SQL.locations"))
		if err != nil {
			panic("failed to read sql file")
		}
		sqlShow, err := os.ReadFile(viper.GetString("SQL.shows"))
		if err != nil {
			panic("failed to read sql file")
		}
		dbNew, errDbNew := common.GetDB()
		if errDbNew != nil {
			log.Panicln(errDbNew)
		}
		dbNew.Exec(string(sqlLocation))
		dbNew.Exec(string(sqlShow))
		config.FirstTime = false
		marshalled, errMarshal := yaml.Marshal(&config)
		if errMarshal != nil {
			panic(fmt.Errorf("error marshalling config, %s", errMarshal))
		}
		errMarshal = os.WriteFile("./config/first.yml", marshalled, 0644)
		if errMarshal != nil {
			panic(fmt.Errorf("error writing yaml file, %s", errMarshal))
		}
		fmt.Println("Configuration has been updated.")
	}

}
