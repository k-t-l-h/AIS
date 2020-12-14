package graphs

import (
	"bufio"
	"github.com/yourbasic/graph"
	"log"
	"os"
	"strings"
)

type GraphContainer struct {
	g *graph.Mutable
	p map[string]int
}

func NewGraphContainer() *GraphContainer {
	return &GraphContainer{}
}

func ( g *GraphContainer) Init() {

	g.p = make(map[string]int)
	file, err := os.Open("graph.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if parts[0] == "node" {
			i++
			g.p[parts[1]] = i
		}
	}


	g.g = graph.New(i+1)
	file2, err := os.Open("graph.txt")
	defer file2.Close()
	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if parts[0] == "edge" {
			g.g.AddBothCost(g.p[parts[1]], g.p[parts[2]], 1)
		}
	}
}

func ( g *GraphContainer) Distance(a, b string) float64{
	i := g.p[a]
	j := g.p[b]

	_, dist :=  graph.ShortestPath(g.g, i, j)
	return float64(dist)
}