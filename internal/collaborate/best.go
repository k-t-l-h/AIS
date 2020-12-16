package collaborate

import (
	"github.com/gonum/stat"
)

func Transpose(a [][]float64) [][]float64 {

	newArr := make([][]float64, len(a[0]))

	for i := 0; i < len(a[0]); i++ {
		newArr[i] = make([]float64, len(a))
		for j := 0; j < len(a); j++ {
			newArr[i][j] = a[j][i]
		}
	}
	return newArr
}

//получаем массив корреляций пользователей
func UserCorrelations(data [][]float64, ind int) []float64 {

	var states []float64
	dataT := Transpose(data)

	ref := dataT[ind]
	for i := 0; i < len(dataT); i++ {
		states = append(states, stat.Correlation(ref, dataT[i], nil))

	}
	return states
}
