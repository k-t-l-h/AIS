package search

import (
	"AIS/internal/data"
	"AIS/internal/models"
	"errors"
)


func Search(strict []bool, a, b models.Article)  {
	Between(a, b)
}

func Between(a, b models.Article)  []models.Article{
	articles := data.Load()
	result := []models.Article{}
	for i := 0; i < len(articles); i++ {
		article := articles[i]

		a, erra := isLess(a, article)
		b, errb := isLess(article, b)

		if erra == nil && errb == nil {
			if a&&b {
				result = append(result, article)
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

	if a.WAK != b.WAK || a.WOS != b.WOS || a.RINZ != b.RINZ {
		return false, errors.New("not compatible")
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