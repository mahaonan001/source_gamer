package utils

import (
	"errors"
	"log"
	"math/rand"
	"regexp"
	"source_gamer/model"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// RandomString 生成随机字符串，用于对chat产生的记录向records_t表种写入提供单独的id
func RandomString(l int, Inner string) string {
	var letters = []byte(Inner)
	var result = make([]byte, l)
	rand.NewSource(time.Now().UnixNano())
	for i := range l {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// row_string 将excel表中的某一格读为字符串返回
func row_string(k int, row []string) string {
	if len(row) < k+1 {
		return ""
	}
	return row[k]
}

// String2int 将字符串类型的数字转为int类型
func String2int(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return n
}

// String2Double 将字符串类型的数字转为double类型
func String2Double(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return -1
	}
	return f
}

// Record 将excel表中的记录写入数据库
func Record(path string, db *gorm.DB) {
	//读取excel表中的记录
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Println(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return
	}
	//舍弃标题行，循环写入数据库
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
			Chat:             false,
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
		var existingRecord model.Record
		// 检查记录是否已存在
		result := db.First(&existingRecord, "id = ?", commen.ID)
		// 如果记录不存在，则创建新记录
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			db.Create(&commen)
		}

	}
}

// Analysis_record 函数用于分析记录
func Analysis_record(path string, db *gorm.DB) {
	// 使用excelize打开文件
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Println(err)
		return
	}
	// 读取excel表中的记录
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return
	}
	// 遍历除第一行外的所有行
	for _, row := range rows[1:] {
		// 如果行中的列数小于18，则跳过此行
		if len(row) < 18 {
			continue
		}
		// 创建一个Dim对象
		var dim model.Dim
		result := db.Where("dim_=?", row[15]).Find(&dim)
		// 在数据库中查询Dim_的值
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			dim.Dim_ = row[15]
			db.Create(&dim)
		}
		// 创建一个Score对象
		score := model.Score{
			RecordId:        row[1],
			Analysis:        row[13],
			Extracted_texts: row[14],
			Dim_id:          dim.Id,
			Option_word:     row[16],
			Score_:          row[17] == "正向",
		}
		var existingScore model.Score
		// 检查记录是否已存在
		results := db.First(&existingScore, "record_id=?", score.RecordId)
		// 如果记录不存在，则创建新记录
		if errors.Is(results.Error, gorm.ErrRecordNotFound) {
			db.Create(&score)
		}
	}
}

// Keyword 函数用于将关键词写入数据库
func Keyword(path string, db *gorm.DB) {
	// 读取excel表中的记录
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Println(err)
		return
	}

	// 读取excel表中的记录
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return
	}
	for _, row := range rows[1:] {
		keyword := model.Keyword{
			RecordId:    row[0],
			T_room:      String2Double(row[10]),
			S_room:      String2int(row_string(11, row)),
			BurnningT:   row_string(12, row),
			HotDevice:   row_string(13, row),
			Device_logo: row_string(14, row),
			Hot_T:       row_string(15, row),
			Time_cyc:    row_string(16, row),
			Money_cyc:   String2Double(row_string(17, row)),
			Gas_cyc:     String2Double(row_string(18, row)),
			Ele_cyc:     String2int(row_string(19, row)),
			Boal_cyc:    String2int(row_string(20, row)),
		}

		var existingKeyword model.Keyword
		// 检查记录是否已存在
		result := db.First(&existingKeyword, "record_id=?", keyword.RecordId)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			db.Create(&keyword)
		}
	}
}
