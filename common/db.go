package common

import (
	"fmt"
	"log"

	"source_gamer/model"

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
		log.Panic("error to connect database,%v", err)
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("error to connect database,%v", err)
	}
	db.AutoMigrate(&model.User{})
	return db
}
func GetDB_User() *gorm.DB {
	return db_User()
}

func db_Commens() *gorm.DB {
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("error to connect database,%v", err)
	}
	db.AutoMigrate(&model.Record{})
	db.Exec("ALTER TABLE records ADD CONSTRAINT fk_profiles_users FOREIGN KEY (email) REFERENCES users(email)")
	return db
}
func GetDB_Commens() *gorm.DB {
	return db_Commens()
}
