package collaborate

import (
	"AIS/internal/models"
	"golang.org/x/exp/errors/fmt"
	"sort"
)


func GetBest(number int, uindex int, all []models.Article, matrix [][]float64) {

	corMatrix := UserCorrelations(matrix, uindex)

	var ac []models.ArticleCorrelation
	for i := 0; i < len(all); i++ {
		ac = append(ac, models.ArticleCorrelation{
			State:   0,
			Article: all[i],
		})
	}

	for i := 0; i < len(ac); i++ {
		for j := 0; j < len(corMatrix); j++ {
			if corMatrix[j] > 0 {
				ac[i].State += matrix[i][j] * corMatrix[j]
			}
		}
	}

	sort.Sort(models.ArticleCorrelations(ac))
	fmt.Printf("Наиболее подходящие вам статьи: ")
	for i := 1; i <= number; i++ {
		fmt.Printf("%+v \n\n", ac[i].Article)
	}

}
