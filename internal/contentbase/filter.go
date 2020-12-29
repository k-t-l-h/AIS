package contentbase

import "github.com/gonum/stat"

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

//средняя оценка пользователя
func UserMedian(data [][]float64) []float64 {

	var states []float64

	data = Transpose(data)

	for i := 0; i < len(data); i++ {
		a := data[i]
		states = append(states, stat.Mean(a, nil))

	}

	return states
}
