package metrix

import (
	"AIS/internal/graphs"
	"AIS/internal/models"
	"github.com/gonum/stat"
	"log"
	"math"
	"reflect"
)

func MakeVector(a, b models.Article) ([]float64, []float64) {
	var vectorA []float64
	var vectorB []float64

	valA := reflect.ValueOf(&a).Elem()
	valB := reflect.ValueOf(&b).Elem()
	for i := 0; i < valA.NumField(); i++ {
		typeField := valA.Type().Field(i).Type.String()

		valueA := valA.Field(i)
		valueB := valB.Field(i)

		switch typeField {
		case "string":
			if valueA.Interface().(string) == valueB.Interface().(string) {
				vectorA = append(vectorA, 1)
				vectorB = append(vectorB, 1)
			} else {
				vectorA = append(vectorA, 1)
				vectorB = append(vectorB, -1)
			}

		case "[]string":

			arrA := valueA.Interface().([]string)
			arrB := valueB.Interface().([]string)
			diff := 0.0
			for i := 0; i < len(arrA) && i < len(arrB); i++ {
				if arrA[i] != arrB[i] {
					diff += 1.0
				}
			}
			vectorA = append(vectorA, diff)
			vectorB = append(vectorB, -diff)

		case "bool":
			if valueA.Interface().(bool)  == valueB.Interface().(bool){
				vectorA = append(vectorA, 1)
				vectorB = append(vectorB, 1)
			} else {
				vectorA = append(vectorA, 1)
				vectorB = append(vectorB, -1)
			}
		case "uint":
			switch valA.Type().Field(i).Name {
			case "Year":
				vectorA = append(vectorA, float64(valueA.Interface().(uint) - 1970)/50)
				vectorB = append(vectorB, float64(valueB.Interface().(uint) - 1970)/50)
			case "Citations":
				vectorA = append(vectorA, float64(valueA.Interface().(uint))/50)
				vectorB = append(vectorB, float64(valueB.Interface().(uint))/50)
			case "ReadingTime":
				vectorA = append(vectorA, float64(valueA.Interface().(uint) + 5)/30)
				vectorB = append(vectorB, float64(valueB.Interface().(uint) + 5)/30)
			}

		case "float64":
			vectorA = append(vectorA, valueA.Interface().(float64))
			vectorB = append(vectorB, valueB.Interface().(float64))
		default:
			log.Print("field is unsupported: ", typeField)
		}
	}

	return vectorA, vectorB
}

func EuclideanDistance(objA, objB models.Article) float64 {

	a, b := MakeVector(objA, objB)
	diff := 0.0
	for i := 0; i < len(a); i++ {
		diff += math.Pow(a[i]-b[i], 2)
	}
	diff = math.Sqrt(diff)

	return diff
}

func CorrelationDistance(objA, objB models.Article) float64 {
	a, b := MakeVector(objA, objB)
	return stat.Correlation(a, b, CorrelationWeight())
}

func GraphDistance(objA, objB models.Article) float64 {
	ct := graphs.NewGraphContainer()
	ct.Init()

	diff := math.Inf(0)
	for i := 0; i < len(objA.Fields); i++ {
		for j := 0; j < len(objB.Fields); j++ {
			diffij := ct.Distance(objA.Fields[i], objB.Fields[j])
			if diffij < diff {
				diff = diffij
			}
		}
	}
	return diff
}

func DiffDistance(objA, objB models.Article) float64 {

	a, b := MakeVector(objA, objB)
	diff := 0.0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff += 1
		}
	}

	return diff
}

func CorrelationWeight()[]float64 {
	//небольшой комментарий
	//наибольший вес у области и оценки
	//наименьшее влияние у года и количества цитат (и названия)
	weight := []float64{0, 0.02, 2, 0.1, 0.1, 0.1, 0.01, 0.01, 1, 0.01}
	return weight
}