package models

type Film struct {
	ID              uint    `json:"ID" gorm:"primary_key"`
	MainGenre       string  `json:"maingenre" gorm:"column:maingenre"`
	RussianName     string  `json:"russianname" gorm:"column:russianname"`
	EnglishName     string  `json:"englishname" gorm:"column:englishname"`
	TrailerLink     string  `json:"trailerlink" gorm:"column:trailerlink"`
	Rating          float64 `json:"rating"`
	ImdbRating      float64 `json:"imdbrating" gorm:"column:imdbrating"`
	TotalVotes      int     `json:"totalvotes" gorm:"column:totalvotes"` //всего голосов
	SumVotes        int     `json:"sumvotes" gorm:"column:sumvotes"`     //сумма голосов,нужна только бэку
	Description     string  `json:"description"`
	Image           string  `json:"image,omitempty"`
	BackgroundImage string  `json:"backgroundimage,omitempty" gorm:"column:backgroundimage"`
	Country         string  `json:"country"`
	Year            int     `json:"year"`
	AgeLimit        int     `json:"agelimit,omitempty" gorm:"column:agelimit"`
}

//easyjson:json
type Films []Film

func (f *Film) TableName() string {
	return "films"
}

type ListFilm struct {
	ID          uint    `json:"ID" gorm:"primary_key"`
	RussianName string  `json:"russianname" gorm:"column:russianname"`
	Image       string  `json:"image,omitempty"`
	Country     string  `json:"country"`
	Year        int     `json:"year"`
	AgeLimit    int     `json:"agelimit,omitempty" gorm:"column:agelimit"`
	Rating      float64 `json:"rating"`
}

//easyjson:json
type ListsFilm []ListFilm

func FormatFilm(film Film) (list ListFilm) {
	list.ID = film.ID
	list.RussianName = film.RussianName
	list.Image = film.Image
	list.Country = film.Country
	list.Year = film.Year
	list.AgeLimit = film.AgeLimit
	list.Rating = film.Rating
	return list
}

func (lists *ListsFilm) Convert(films Films) ListsFilm {
	var lf ListsFilm
	for _, f := range films {
		lf = append(lf, FormatFilm(f))
	}
	return lf
}

//func (film *Film) SetRating(rating float64) {
//	film.Rating = rating
//}
