package contentbase

import (
	"AIS/internal/metrix"
	"AIS/internal/models"
	"fmt"
	"sort"
)

func GetBest(number int, a models.Article, all []models.Article) {

	var ac []models.ArticleCorrelation
	for i := 0; i < len(all); i++ {
		ac = append(ac, models.ArticleCorrelation{
			State:   0,
			Article: all[i],
		})
	}

	for i := 0; i < len(ac); i++ {
		ac[i].State = metrix.CorrelationDistance(a, ac[i].Article)
	}

	sort.Sort(models.ArticleCorrelations(ac))
	fmt.Printf("Наиболее подходящие вам статьи: ")
	for i := 1; i <= number; i++ {
		fmt.Printf("%+v \n\n", ac[i].Article)
	}

}
