package contentbase

import (
	"AIS/internal/models"
	"fmt"
	"github.com/gonum/stat"
	"sort"
)


func GetBest(number int, index int, all []models.Article, matrix [][]float64) {

	corMatrix := UserMedian(matrix)

	var ac []models.ArticleCorrelation
	for i := 0; i < len(all); i++ {
		ac = append(ac, models.ArticleCorrelation{
			Index:   i,
			State:   0,
			Article: all[i],
		})
	}

	ref := matrix[index]
	for i := 0; i < len(ac); i++ {
		var refc []float64
		for j := 0; j < len(matrix[0]); j++ {
			refc = append(refc, matrix[i][j] - corMatrix[j])
		}
		ac[i].State = stat.Correlation(ref, refc, nil)
	}


	sort.Sort(models.ArticleCorrelations(ac))
	fmt.Printf("Наиболее подходящие вам статьи: \n")

	k := 0
	for i := 0; k < number && i < len(ac) && ac[i].State > 0; i++ {
		if ac[i].Index != index {
			k++
			fmt.Printf("Счёт: %f Cтатья: %s \n\n", ac[i].State, ac[i].Article.Title)
		}
	}


}