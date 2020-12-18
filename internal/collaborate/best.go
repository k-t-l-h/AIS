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
		a, b := findUnion(ref, dataT[i])
		states = append(states, stat.Correlation(a, b, nil))

	}
	return states
}

func findUnion(ref, data []float64) ([]float64, []float64) {

	var stateRef []float64
	var stateData []float64

	for i := 0; i < len(ref); i++ {
		if ref[i] != 0 && data[i] != 0 {
			stateRef = append(stateRef, ref[i])
			stateData = append(stateData, data[i])
		}
	}
	return stateRef, stateData
}
