package search

import (
	"AIS/internal/data"
	"AIS/internal/metrix"
	"AIS/internal/models"
	"errors"
	"github.com/gonum/stat"
	"log"
	"sort"
)


func Search(strict []bool, a, b models.Article)  {
	_a, _b := a, b
	res := Between(a, b)

	v := 0
	for len(res) != 0 || v < 5 {
		v++
		a, b = Expand(strict, a, b)
		res = Between(a, b)
	}


	for i := 0; i < len(res); i++ {
		a, sa := metrix.MakeVector(_a, res[i].Article)
		res[i].State += stat.Correlation(a, sa, nil)
		b, sb := metrix.MakeVector(_b, res[i].Article)
		res[i].State += stat.Correlation(b, sb, nil)
		res[i].State /= 2
	}

	sort.Sort(models.ArticleCorrelations(res))

	for i := 0; i < len(res); i++ {
		log.Printf("%+v\n", res[i].Article)
	}
}

func Between(a, b models.Article)  []models.ArticleCorrelation{
	articles := data.Load()
	result := []models.ArticleCorrelation{}

	for i := 0; i < len(articles); i++ {
		article := articles[i]

		a, erra := isLess(a, article)
		b, errb := isLess(article, b)

		if erra == nil && errb == nil {
			if a&&b {
				result = append(result,
					models.ArticleCorrelation{Article: article})
			}
		}
	}
	return result
}

func isLess(a, b models.Article)  (bool, error){

	//если области не совпали, это ошибка
	flag := false
	for i := 0; i < len(a.Fields) ; i++ {
		for j := 0; j < len(b.Fields); j++ {
			if a.Fields[i] == b.Fields[j]{
				flag = true
			}
		}
	}
	if !flag {
		return flag, errors.New("not compatible")
	}

	if !(a.WAK&&b.WAK || a.WOS&&b.WOS || a.RINZ&&b.RINZ) {
		return false, nil
	}

	if a.Year > b.Year {
		return false, nil
	}

	if a.Citations > b.Citations {
		return false, nil
	}

	if a.Score > b.Score {
		return false, nil
	}

	if a.ReadingTime > b.ReadingTime {
		return false, nil
	}

	return true, nil
}