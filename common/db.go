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

func Init_db() error {
	err := easyX.CreateDB(viper.GetString("datasource.database"),
		viper.GetString("datasource.username"),
		viper.GetString("datasource.password"),
		viper.GetString("datasource.hostname"),
		viper.GetInt("datasource.port"),
	)
	return err
}

func getDB() (*gorm.DB, error) {
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
		log.Printf("error to connect database,%v\n", err)
		return nil, err
	}
	err = db.AutoMigrate(&model.EmailCode{}, &model.User{}, &model.Record{}, &model.Chat{}, &model.Score{}, &model.Dim{}, &model.Keyword{}, &model.Location{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDB() (*gorm.DB, error) {
	return getDB()
}
