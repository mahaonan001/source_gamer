package utils

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Timely(db *gorm.DB) {
	Record(viper.GetString("WorkDir.record"), db)

	Analysis_record(viper.GetString("WorkDir.score"), db)

	Keyword(viper.GetString("WorkDir.keyword"), db)
}
