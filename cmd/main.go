package main

import (
	"AIS/internal/metrix"
	"AIS/internal/models"
	"log"
)

func main() {
	a := models.Article{
		Title:       "A",
		Authors:     []string{"math"},
		Fields:       []string{"pure_math", "applied_math"},
		RINZ:        true,
		WAK:         false,
		WOS:         false,
		Year:        1999,
		Citations:   0,
		Score:       5,
		ReadingTime: 10,
	}

	b := models.Article{
		Title:       "A",
		Authors:     []string{"A", "B"},
		Fields:       []string{"applied_math"},
		RINZ:        true,
		WAK:         false,
		WOS:         false,
		Year:        1999,
		Citations:   0,
		Score:       5,
		ReadingTime: 10,
	}

	log.Print(metrix.CorrelationDistance(a, b))
}
