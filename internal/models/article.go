package models

type Article struct {
	//номинальные признаки
	Title string
	Authors []string
	Fields []string

	//бинарные признаки
	RINZ bool
	WAK bool
	WOS bool

	//вещественные признаки
	Year uint
	Citations uint
	Score float64
	ReadingTime uint
}
