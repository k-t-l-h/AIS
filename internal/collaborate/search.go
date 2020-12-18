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
			Index:   i,

			State:   0,
			Article: all[i],
		})
	}

	for i := 0; i < len(ac); i++ {
		for j := 0; j < len(corMatrix); j++ {
			if corMatrix[j] > 0 && matrix[i][j] > 0 {
				ac[i].State += matrix[i][j] * corMatrix[j]
			}
		}
	}


	sort.Sort(models.ArticleCorrelations(ac))
	fmt.Printf("Наиболее подходящие вам статьи: \n")
	k := 0
	for i := 0; k <= number || i < len(ac); i++ {
		if matrix[ac[i].Index][uindex] == 0 {
			k++
			fmt.Printf("Счёт: %f Cтатья: %s \n\n", ac[i].State, ac[i].Article.Title)
		}
	}

}
