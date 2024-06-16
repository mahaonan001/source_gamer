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
	//使用单独的路由管理文件
	r = router.CollectRouter(r)
	err := r.Run(":" + viper.GetString("server.port"))
	if err != nil {
		return
	}
}
func init() { //init函数会在main前执行，用于初始化任务
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	} //读取yaml文件
	viper2 := viper.New()
	viper2.SetConfigName("first")
	viper2.SetConfigType("yml")
	viper2.AddConfigPath(workDir + "/config") //读取first.yml
	err2 := viper2.ReadInConfig()
	if err2 != nil {
		panic(err2)
	}
	var config model.Config
	err2 = viper2.UnmarshalExact(&config) //将first.yml文件的参数绑定到结构体上
	if err2 != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err2))
	}
	if config.FirstTime { //判断是否进行初始化
		errInit := common.Init_db() //初始化数据库，创建需要的数据库和表
		if errInit != nil {
			log.Panicln("err", errInit)
			return
		}
		//执行读取sql文件位置为sql语句执行做准备
		sqlLocation, err := os.ReadFile(viper.GetString("SQL.locations"))
		if err != nil {
			panic("failed to read sql file")
		}
		sqlShow, err := os.ReadFile(viper.GetString("SQL.shows"))
		if err != nil {
			panic("failed to read sql file")
		}
		//获取初始化后的数据库
		dbNew, errDbNew := common.GetDB()
		if errDbNew != nil {
			log.Panicln(errDbNew)
		}
		//执行sql语句创建视图
		dbNew.Exec(string(sqlLocation))
		dbNew.Exec(string(sqlShow))
		//初始化完成，将first.yml参数修改并写回原文件
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
