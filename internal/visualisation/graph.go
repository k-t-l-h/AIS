package visualisation

import (
	"bufio"
	"github.com/go-echarts/go-echarts/charts"
	"golang.org/x/exp/rand"
	"log"
	"os"
	"strings"
)

func testgraph()  {
	var graphNodes []charts.GraphNode
	links := make([]charts.GraphLink, 0)

	file, err := os.Open("graph.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		line := scanner.Text()
		i++
		parts := strings.Split(line, " ")

		if parts[0] == "node" {
			size := 10*rand.Float32()
			graphNodes = append(graphNodes, charts.GraphNode{Name: parts[1], SymbolSize: size})
		}

		if parts[0] == "edge" {
			links = append(links, charts.GraphLink{Source: parts[1], Target: parts[2]})
		}
	}

	graph := charts.NewGraph()
	graph.Add("Graph", graphNodes, links,
		charts.GraphOpts{Roam: true, FocusNodeAdjacency: true, Force: charts.GraphForce{
			Repulsion:  8000,
		}},
		charts.EmphasisOpts{Label: charts.LabelTextOpts{Show: true, Position: "left", Color: "black"},
			ItemStyle: charts.ItemStyleOpts{Color: "yellow"}},
		charts.LineStyleOpts{Curveness: 0.2})
	f, _ := os.Create("graph.html")
	graph.Render(f)

}
