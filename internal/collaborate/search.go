package collaborate

import (
	"AIS/internal/models"
	"golang.org/x/exp/errors/fmt"
	"sort"
)

type ArticleCorrelation struct {
	State   float64
	Article models.Article
}

type ArticleCorrelations []ArticleCorrelation

func (a ArticleCorrelations) Len() int           { return len(a) }
func (a ArticleCorrelations) Less(i, j int) bool { return a[i].State > a[j].State }
func (a ArticleCorrelations) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetBest(number int, uindex int, all []models.Article, matrix [][]float64) {

	corMatrix := UserCorrelations(matrix, uindex)

	var ac []ArticleCorrelation
	for i := 0; i < len(all); i++ {
		ac = append(ac, ArticleCorrelation{
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

	sort.Sort(ArticleCorrelations(ac))
	fmt.Printf("Наиболее подходящие вам статьи: ")
	for i := 1; i <= number; i++ {
		fmt.Printf("%+v \n\n", ac[i].Article)
	}

}
