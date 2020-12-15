package main

import (
	"AIS/internal/data"
	"AIS/internal/metrix"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите название статьи-эталона: ")
	A, _ := reader.ReadString('\n')
	A = strings.Replace(A, "\n", "", -1)
	fmt.Println("---------------------")
	fmt.Println("Введите название второй статьи: ")
	B, _ := reader.ReadString('\n')
	B = strings.Replace(B, "\n", "", -1)
	fmt.Println("---------------------")

	articles := data.Load()
	a, erra := data.Find(articles, A)
	b, errb := data.Find(articles, B)
	if erra != nil || errb != nil {
		log.Print(erra, errb)
		fmt.Println("Одна из статей не найдена :( ")
		os.Exit(2)
	}

	fmt.Printf("Статья А: %+v \n", a)
	fmt.Printf("Статья B: %+v \n", b)
	fmt.Println("Значения метрик: ")
	fmt.Println("Евклидово расстояние: ", metrix.EuclideanDistance(a, b))
	fmt.Println("Корреляция: ", metrix.CorrelationDistance(a, b))
	fmt.Println("Расстояние по графу: ", metrix.GraphDistance(a, b))
	fmt.Println("Количество различий: ", metrix.DiffDistance(a, b))
}
