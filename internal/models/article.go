package models

type Article struct {
	//номинальные признаки
	Title   string   `json:"title"`
	Authors string `json:"authors"`
	Fields  string `json:"fields"`

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

type ArticleCorrelation struct {
	State   float64
	Article Article
}

type ArticleCorrelations []ArticleCorrelation

func (a ArticleCorrelations) Len() int           { return len(a) }
func (a ArticleCorrelations) Less(i, j int) bool { return a[i].State > a[j].State }
func (a ArticleCorrelations) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
