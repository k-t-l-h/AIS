package main

import (
	"bufio"
	"github.com/go-echarts/go-echarts/charts"
	"log"
	"os"
	"strings"
)



func main()  {

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
			graphNodes = append(graphNodes, charts.GraphNode{Name: parts[1]})
		}

		if parts[0] == "edge" {
			links = append(links, charts.GraphLink{Source: parts[1], Target: parts[2]})
		}
	}

	graph := charts.NewGraph()
	graph.Add("Graph", graphNodes, links,
		charts.GraphOpts{Layout: "circular", Force: charts.GraphForce{Repulsion: 8000}})
	f, _ := os.Create("graph.html")
	graph.Render(f)
}