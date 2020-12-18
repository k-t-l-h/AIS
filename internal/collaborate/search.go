package collaborate

import (
	"AIS/internal/models"
	"golang.org/x/exp/errors/fmt"
	"sort"
)

type ArticleCorrelation struct {
	Index   int
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
			Index:   i,
			State:   0,
			Article: all[i],
		})
	}

	for i := 0; i < len(ac); i++ {
		k := 0.0
		for j := 0; j < len(corMatrix); j++ {
			if corMatrix[j] > 0 && matrix[i][j] > 0 {
				k+=1
				ac[i].State += matrix[i][j] * corMatrix[j]
			}
		}
		ac[i].State /= k
	}

	sort.Sort(ArticleCorrelations(ac))
	fmt.Printf("Наиболее подходящие вам статьи: \n")
	k := 0
	for i := 0; k < number && i < len(ac); i++ {
		if matrix[ac[i].Index][uindex] == 0 {
			k++
			fmt.Printf("Счёт: %f Cтатья: %s \n\n", ac[i].State, ac[i].Article.Title)
		}
	}

}
