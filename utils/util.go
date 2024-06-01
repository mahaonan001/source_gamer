package utils

import (
	"log"
	"math/rand"
	"regexp"
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

func Test_xslx(path, email string) {
	var commens []model.Record
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
			Email:            email,
			V_type:           row[0],
			Coding:           row[1],
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
}
