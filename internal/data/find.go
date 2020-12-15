package data

import (
	"AIS/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Load() []models.Article {
	jsonFile, err := os.Open("data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var obj []models.Article
	json.Unmarshal(byteValue, &obj)
	return obj
}

func Find(articles []models.Article, name string) (models.Article, error) {
	for i := 0; i < len(articles); i++ {
		if strings.ToLower(articles[i].Title) == strings.ToLower(name) {
			return articles[i], nil
		}
	}
	return models.Article{}, errors.New("not found")
}
