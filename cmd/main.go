package main

import (
	"AIS/internal/data"
	"AIS/internal/search"
)

func main() {

	a, b := data.LoadParam()
	keys := data.LoadKeys()

	search.Search(keys, a, b)
}
