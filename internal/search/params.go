package search

import (
	"AIS/internal/data"
	"AIS/internal/metrix"
	"AIS/internal/models"
	"errors"
	"fmt"
	"github.com/gonum/stat"
	"sort"
)


func Search(strict []bool, a, b models.Article) {
	_a, _b := a, b
	res := Between(a, b)

	v := 0
	if len(res) == 0 {

		for ; v < 5 ; v++ {
			a, b = Expand(strict, a, b)
			res = Between(a, b)

			if len(res) > 0 {
				break
			}
		}
	}
	if len(res) == 0 {
		fmt.Printf("Рекомендации найти не удалось \n")
		return
	}
	if v == 0 {
		fmt.Printf("Рекомендации найдены за первую итерацию \n")
	} else {
		fmt.Printf("Рекомендации найдены за %d итераций \n", v)
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
		fmt.Printf("%+v\n", res[i].Article)
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
				//log.Print(a.Fields[i])
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