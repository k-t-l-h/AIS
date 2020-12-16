package models

type Article struct {
	//номинальные признаки
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Fields  []string `json:"fields"`

	//бинарные признаки
	RINZ bool `json:"rinz"`
	WAK  bool `json:"wak"`
	WOS  bool `json:"wos"`

	//вещественные признаки
	Year        uint    `json:"year"`
	Citations   uint    `json:"citations"`
	Score       float64 `json:"score"`
	ReadingTime uint    `json:"reading_time"`
}
