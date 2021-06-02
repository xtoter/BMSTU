package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

type Vertex struct {
	old int
	new int
}

var time int
var used []bool

type data struct {
	a       []int
	b       []rune
	n, m, q int
	order   []int
	V       []*Vertex
	G       [][]*Vertex
}

func dfs(data data) []int {
	return VisitVertex(data.V[data.q], data)
}

func VisitVertex(v *Vertex, data data) []int {
	v.new = time
	time++
	used[v.old] = true
	data.order = append(data.order, v.old)
	for i := 0; i < len(data.G[v.old]); i++ {
		u := data.G[v.old][i]
		if !used[u.old] {
			data.order = VisitVertex(u, data)
		}
	}
	return data.order
}
func new(t []int) []int {
	new1 := make([]int, len(t))
	for i := 0; i < len(t); i++ {
		new1[t[i]] = i
	}
	return new1
}
func print(data data) {
	fmt.Printf("%d\n%d\n%d\n", len(data.order), data.m, data.V[data.q].new)
	var j int
	for m := 0; m < len(data.order); m++ {
		if used[data.V[data.order[m]].old] {
			for j = 0; j < data.m; j++ {
				fmt.Printf("%d ", data.V[data.a[data.order[m]*data.m+j]].new)
			}
			fmt.Printf("\n")
		}
	}
	for m := 0; m < len(data.order); m++ {
		if used[data.V[data.order[m]].old] {
			for j = 0; j < data.m; j++ {
				fmt.Printf("%c ", data.b[data.order[m]*data.m+j])
			}
			fmt.Printf("\n")
		}
	}

}

func main() {
	time = 0
	var data data
	var temp1 int
	var temp2 rune

	input.Scanf("%d%d%d", &data.n, &data.m, &data.q)
	data.V = make([]*Vertex, data.n)
	data.G = make([][]*Vertex, data.n)
	for i := 0; i < data.n; i++ {
		var temp Vertex
		temp.old = i
		var t bool
		t = false
		used = append(used, t)
		data.V[i] = &temp
	}
	for i := 0; i < data.n*data.m; i++ {
		input.Scanf("%d", &temp1)
		data.a = append(data.a, temp1)
		data.G[i/data.m] = append(data.G[i/data.m], data.V[temp1])
	}
	//fmt.Print(G)
	i := 0
	for i < data.n*data.m {
		input.Scanf("%c", &temp2)
		if temp2 != ' ' && temp2 != '\n' {

			data.b = append(data.b, temp2)
			i++
		}
	}
	data.order = dfs(data)
	print(data)
}

