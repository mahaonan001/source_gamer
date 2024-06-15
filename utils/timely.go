package utils

import "github.com/spf13/viper"

func Timely() {
	Record(viper.GetString("WorkDir.record"))

	Analysis_record(viper.GetString("WorkDir.score"))

	Keyword(viper.GetString("WorkDir.keyword"))
}
