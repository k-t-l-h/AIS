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

func LoadParam() (models.Article, models.Article) {
	jsonFile, err := os.Open("params.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var obj []models.Article
	json.Unmarshal(byteValue, &obj)

	if len(obj) < 2 {
		fmt.Println(errors.New("not enough data"))
		os.Exit(1)
	}
	return obj[0], obj[1]
}


func LoadKeys() ([]bool) {
	jsonFile, err := os.Open("keys.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var obj []bool
	json.Unmarshal(byteValue, &obj)

	return obj
}

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

func FindUser(users []string, name string) (int, error) {
	for i := 0; i < len(users); i++ {
		if strings.ToLower(users[i]) == strings.ToLower(name) {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func FindArticles(articles []models.Article, name string) (int, error) {
	for i := 0; i < len(name); i++ {
		if strings.ToLower(articles[i].Title) == strings.ToLower(name) {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}