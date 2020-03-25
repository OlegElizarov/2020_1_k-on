package models

type Series struct {
	ID       uint   `json:"ID" gorm:"primary_key"`
	SeasonId uint   `json:"seasonid" gorm:"column:seasonid"`
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Image    string `json:"image,omitempty"`
}

//easyjson:json
type SeriesArr []Series

func (s *Series) TableName() string {
	return "series"
}

//триггер на бд авто инсерт инто тейблс рейтингс
//общ рейтинг и кол-во рейтингов
//цепь
