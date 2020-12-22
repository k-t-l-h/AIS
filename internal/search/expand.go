package search

import (
	"AIS/internal/graphs"
	"AIS/internal/models"
)

func Expand(params []bool, a, b models.Article) (models.Article, models.Article){

	if params[2] {
		a, b = expandField(a, b)
	}
	if params[3] {
		a, b = expandRinz(a, b)
	}
	if params[4] {
		a, b = expandWak(a, b)
	}
	if params[5] {
		a, b = expandWos(a, b)
	}
	if params[6] {
		a, b = expandYear(a, b)
	}
	if params[7] {
		a, b = expandCitations(a, b)
	}
	if params[8] {
		a, b = expandScore(a, b)
	}
	if params[9] {
		a, b = expandReadingTime(a, b)
	}

	return a, b
}

func expandField(a, b models.Article) (models.Article, models.Article){
	g := graphs.NewGraphContainer()
	g.Init()

	var s []string
	for i := 0; i < len(a.Fields); i++ {
		s = append(s,g.GetNearest(a.Fields[i])...)
	}
	a.Fields = s

	var s2 []string
	for i := 0; i < len(b.Fields); i++ {
		s2 = append(s2,g.GetNearest(b.Fields[i])...)
	}
	b.Fields = s2
	return a, b
}

func expandRinz(a, b models.Article) (models.Article, models.Article){
	a.RINZ = false
	b.RINZ = true
	return a, b
}

func expandWos(a, b models.Article)(models.Article, models.Article) {
	a.WOS = false
	b.WOS = true
	return a, b
}

func expandWak(a, b models.Article) (models.Article, models.Article){
	a.WAK = false
	b.WAK = true
	return a, b
}

func expandYear(a, b models.Article)(models.Article, models.Article) {
	a.Year -= 1
	b.Year += 1
	return a, b
}

func expandCitations(a, b models.Article) (models.Article, models.Article){
	a.Citations -= 1
	b.Citations += 1
	return a, b
}

func expandScore(a, b models.Article) (models.Article, models.Article){
	a.Score -= 0.1
	b.Score += 0.1
	return a, b
}

func expandReadingTime(a, b models.Article)(models.Article, models.Article) {
	a.ReadingTime -= 1
	b.ReadingTime += 1
	return a, b
}