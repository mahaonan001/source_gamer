package common

import (
	"fmt"
	"net/http"
	"source_gamer/response"

	"source_gamer/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	easyX "github.com/mahaonan001/easyX"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init_db() {
	easyX.CreateDB(viper.GetString("datasource.database"),
		viper.GetString("datasource.username"),
		viper.GetString("datasource.password"),
		viper.GetInt("datasource.port"),
	)
}

func code_email_DB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		var c *gin.Context
		response.Response(c, http.StatusServiceUnavailable, 500, gin.H{"code": 500}, "数据库连接出错")
		panic(err)
	}
	db.AutoMigrate(&model.EmailCode{})
	return db
}

func GetDB_Email() *gorm.DB {
	return code_email_DB()
}
func db_User() *gorm.DB {
	host := viper.GetString("datasource.hostname")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	var c *gin.Context
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		response.Response(c, http.StatusServiceUnavailable, 500, gin.H{"code": 500}, "数据库连接出错")
	}
	db.AutoMigrate(&model.User{})
	return db
}
func GetDB_User() *gorm.DB {
	return db_User()
}
