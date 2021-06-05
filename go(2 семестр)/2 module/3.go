package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

type Component struct {
	vertred   map[int]bool
	vertices  int
	vertmin   int
	neighbors int
}

type Vertex struct {
	used   bool
	advert []int
}
type Graph struct {
	components []Component
	tops       []Vertex
	compmax    int
}

func findingcomponent(graph Graph) Graph {
	graph.compmax = 0
	for i := 1; i < len(graph.components); i++ {
		if compare(graph, i) {
			graph.compmax = i
		}
	}
	return graph
}
func color(graph Graph, n int) Graph {
	graph.components = make([]Component, 0)
	for i, j := 0, 0; i < n; i++ {
		if !graph.tops[i].used {
			var comp Component
			comp.vertices = 0
			comp.neighbors = 0
			comp.vertmin = n
			comp.vertred = make(map[int]bool)
			graph.components = append(graph.components, comp)
			dfs(i, j, graph)
			j++
		}
	}
	return graph
}
func dfs(v, id int, graph Graph) {
	graph.tops[v].used = true
	graph.components[id].vertices++
	graph.components[id].vertred[v] = true
	if graph.components[id].vertmin > v {
		graph.components[id].vertmin = v
	}
	for i := 0; i < len(graph.tops[v].advert); i++ {
		graph.components[id].neighbors++
		temp := graph.tops[v].advert[i]
		if !graph.tops[temp].used {
			dfs(temp, id, graph)

		}
	}
}
func initvertex(n int) []Vertex {
	var graph []Vertex
	for i := 0; i < n; i++ {
		var temp Vertex
		temp.advert = make([]int, 0)
		temp.used = false
		graph = append(graph, temp)
	}
	return graph

}
func compare(graph Graph, b int) bool {
	a := graph.compmax
	if graph.components[a].vertices != graph.components[b].vertices {
		return graph.components[a].vertices < graph.components[b].vertices
	}
	if graph.components[a].neighbors != graph.components[b].neighbors {
		return graph.components[a].neighbors < graph.components[b].neighbors
	}
	return graph.components[a].vertmin > graph.components[b].vertmin
}
func print(graph Graph) {
	fmt.Print("graph {\n")
	for i := 0; i < len(graph.tops); i++ {
		if graph.components[graph.compmax].vertred[i] {

			fmt.Print(i, " [color = red]\n")
		} else {
			fmt.Print(i, "\n")
		}
	}
	for i := 0; i < len(graph.tops); i++ {
		for j := 0; j < len(graph.tops[i].advert); j++ {
			to := graph.tops[i].advert[j]
			if i <= to {
				if graph.components[graph.compmax].vertred[i] {
					fmt.Print(i, " -- ", to, " [color = red]\n")
				} else {
					fmt.Print(i, " -- ", to, "\n")
				}
			}
		}
	}
	fmt.Print("}")
}
func main() {
	var graph Graph
	var n, m, temp1, temp2 int
	input.Scanf("%d%d", &n, &m)
	graph.tops = initvertex(n)
	for i := 0; i < m; i++ {
		input.Scanf("%d%d", &temp1, &temp2)
		if !(temp1 == temp2) {
			graph.tops[temp1].advert = append(graph.tops[temp1].advert, temp2)
		}
		graph.tops[temp2].advert = append(graph.tops[temp2].advert, temp1)
	}
	graph = color(graph, n)
	graph = findingcomponent(graph)
	print(graph)
}

