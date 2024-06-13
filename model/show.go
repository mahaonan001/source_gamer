package model

type Show struct {
	ID               string
	Cleaned_comments string
	Option_word      string
	Score_           bool
	T_room           float64
	S_room           int
	BurnningT        string
	Device_logo      string
	Hot_T            string
	Time_cyc         string
	Money_cyc        float64
	Gas_cyc          float64
	Ele_cyc          int
	Boal_cyc         int
}

func (s Show) TableName() string {
	return "view"
}
