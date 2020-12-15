package data

import (
	"AIS/internal/graphs"
	"AIS/internal/models"
	"encoding/json"
	"github.com/bxcodec/faker"
	"io/ioutil"
	"math/rand"
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
		obj.Authors = make([]string, rand.Intn(3)+1)
		for i := 0; i < len(obj.Authors); i++ {
			obj.Authors[i] = faker.Name()
		}

		obj.Fields = make([]string, 1)
		obj.Fields[0] = fields[rand.Intn(len(fields))]
		obj.Citations = uint(rand.Intn(41))
		obj.ReadingTime = uint(rand.Intn(25) + 5)
		objs[i] = obj
		names[i] = obj.Title
	}

	file, _ := json.MarshalIndent(objs, "", "")

	_ = ioutil.WriteFile("data.json", file, 0644)

	name, _ := json.MarshalIndent(names, "", "")

	_ = ioutil.WriteFile("names.json", name, 0644)
}
