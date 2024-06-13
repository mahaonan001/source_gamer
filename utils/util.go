package utils

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"source_gamer/common"
	"source_gamer/model"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

func RandomString(l int, Inner string) string {
	var letters = []byte(Inner)
	var result = make([]byte, l)
	rand.NewSource(time.Now().UnixNano())
	for i := range l {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func Test_xslx(path string) {
	var commens []model.Record
	db := common.GetDB()
	f, err := excelize.OpenFile("a.xlsx")

	if err != nil {
		log.Println(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return
	}
	for _, row := range rows[1:] {
		pg_n, err := strconv.ParseInt(row[3], 10, 64)
		if err != nil {
			return
		}
		pattern := `\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}`
		re, err := regexp.Compile(pattern)
		if err != nil {
			log.Fatal(err)
		}
		matches := re.FindString(row[7])
		ln, err := strconv.ParseInt(row[9], 10, 64)
		if err != nil {
			return
		}
		commen := model.Record{
			V_type:           row[0],
			ID:               row[1],
			V_link:           row[2],
			Page_n:           pg_n,
			User_name:        row[4],
			User_id:          row[5],
			User_home:        row[6],
			Time:             matches,
			Ip:               row[8],
			Like_n:           ln,
			Like_l:           row[10],
			Cleaned_comments: row[11],
		}
		commens = append(commens, commen)
	}

	db.Create(&commens)
	if db.Error != nil {
		return
	}
}

func Analysis_record(path string) {
	// var commens []model.Record
	f, err := excelize.OpenFile("b.xlsx")
	db := common.GetDB()
	if err != nil {
		log.Println(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return
	}

	for _, row := range rows[1:] {

		if len(row) < 18 {
			continue
		}
		var dim model.Dim
		// var score model.Score
		result := db.Where("dim_=?", row[15]).Find(&dim)
		if result.Error != nil {
			fmt.Println("Error occurred during querying:", result.Error)
		} else if result.RowsAffected == 0 {
			dim.Dim_ = row[15]
			db.Create(&dim)
		}
		break

	}
	// db := common.GetDB()
	// db.Create(&commens)
	// if db.Error != nil {
	// 	return
	// }
}
