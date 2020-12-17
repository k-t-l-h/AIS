package data

import (
	"AIS/internal/graphs"
	"AIS/internal/models"
	"encoding/csv"
	"fmt"
	"github.com/bxcodec/faker"
	"math/rand"
	"os"
)

func Generate() {
	g := graphs.NewGraphContainer()
	fields := g.Fields()

	names := make([]string, 100)
	objs := make([]models.Article, 100)
	for i := 0; i < len(objs); i++ {

		obj := objs[i]
		obj.Score = rand.Float64()
		obj.Year = uint(rand.Intn(50) + 1970)
		typ := rand.Intn(3)
		switch typ {
		case 0:
			obj.RINZ = true
		case 1:
			obj.WAK = true
		default:
			obj.WOS = true
		}
		obj.Title = faker.Sentence()
		obj.Authors = faker.Name()


		obj.Fields = fields[rand.Intn(len(fields))]
		obj.Citations = uint(rand.Intn(41))
		obj.ReadingTime = uint(rand.Intn(25) + 5)
		objs[i] = obj
		names[i] = obj.Title
		fmt.Printf("('%v', '%v', '%v', %v, %v, %v, %v, %v, %v, %v),", obj.Title, obj.Authors, obj.Fields,
			obj.RINZ, obj.WAK, obj.WOS, obj.Year, obj.Citations, obj.Score, obj.ReadingTime)
	}


}

func GenerateUserData() {
	file, _ := os.Create("users.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var data [][]string
	data = make([][]string, 101)
	data[0] = make([]string, 101)
	data[0][0] = "Статьи"
	for i := 0; i < 100; i++ {
		name := faker.Name()
		data[0][i+1] = name
	}

	art := Load()
	for i := 0; i < 100; i++ {
		data[i+1] = make([]string, 101)
		data[i+1][0] = art[i].Title
		for j := 0; j < 100; j++ {
			data[i+1][j+1] = fmt.Sprintf("%d", rand.Intn(5))
		}
	}

	for _, value := range data {
		writer.Write(value)
	}

}
