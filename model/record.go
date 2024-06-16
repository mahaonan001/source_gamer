package model

type Record struct {
	ID               string `gorm:"type:varchar(10);primaryKey;not null"`
	Chat             bool   `gorm:"type:bool;not null"`
	V_type           string `gorm:"type:varchar(10);"`
	V_link           string `gorm:"type:varchar(100);"`
	Page_n           int64  `gorm:"type:bigint;"`
	User_name        string `gorm:"type:varchar(20)"`
	User_id          string `gorm:"type:varchar(20)"`
	User_home        string `gorm:"type:varchar(255)"`
	Time             string `gorm:"type:varchar(40)"`
	Ip               string `gorm:"type:varchar(10)"`
	Like_n           int64  `gorm:"type:bigint"`
	Like_l           string `gorm:"type:varchar(5)"`
	Cleaned_comments string `gorm:"type:longtext;not null"`
}

func (r *Record) TableName() string {
	return "records_t"
}

type Score struct {
	RecordId        string `gorm:"type:index;not null"`
	Record          Record `gorm:"foreignKey:RecordId"`
	Analysis        string `gorm:"type:longtext;not null"`
	Extracted_texts string `gorm:"type:longtext;not null"`
	Dim_id          int    `gorm:"type:index;not null"`
	Dim             Dim    `gorm:"foreignKey:Dim_id"`
	Option_word     string `gorm:"type:varchar(55);not null"`
	Score_          bool   `gorm:"type:int;not null"`
}

func (s *Score) TableName() string {
	return "scores_t"
}

type Keyword struct {
	RecordId    string  `gorm:"type:index;not null;1"`
	Record      Record  `gorm:"foreignKey:RecordId"`
	T_room      float64 `gorm:"type:double;"`
	S_room      int     `gorm:"type:int"`
	BurnningT   string  `gorm:"type:varchar(20)"`
	Device_logo string  `gorm:"type:varchar(20)"`
	Hot_T       string  `gorm:"type:varchar(20)"`
	Time_cyc    string  `gorm:"type:varchar(20)"`
	Money_cyc   float64 `gorm:"type:double"`
	Gas_cyc     float64 `gorm:"type:double"`
	Ele_cyc     int     `gorm:"type:int"`
	Boal_cyc    int     `gorm:"type:int"`
}

func (k *Keyword) TableName() string {
	return "keywords_t"
}

type Dim struct {
	Id   int    `gorm:"type:int;autoincrement;primaryKey;not null"`
	Dim_ string `gorm:"type:varchar(10);not null;unique"`
}

func (d *Dim) TableName() string {
	return "dims_t"
}

type Chat struct {
	EmailId  string `gorm:"type:index;not null"`
	User     User   `gorm:"foreignKey:EmailId;references:Email"`
	RecordId string `gorm:"index"`
	Record   Record `gorm:"foreignKey:RecordId;references:ID"`
}

func (c *Chat) TableName() string {
	return "chats_t"
}
