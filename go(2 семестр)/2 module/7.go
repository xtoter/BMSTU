package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

func Prim(n int, g [][]int) int {
	var result int
	var used []bool
	var min_e, sel_e []int
	for i := 0; i < n; i++ {
		used = append(used, false)
	}
	for i := 0; i < n; i++ {
		min_e = append(min_e, 10000)
	}
	for i := 0; i < n; i++ {
		sel_e = append(sel_e, -1)
	}
	min_e[0] = 0
	for i := 0; i < n; i++ {
		v := -1
		for j := 0; j < n; j++ {
			if !used[j] && (v == -1 || min_e[j] < min_e[v]) {
				v = j
			}
		}
		if min_e[v] == 10000 {
		}
		used[v] = true
		if sel_e[v] != -1 {
			result = result + g[v][sel_e[v]]
			//fmt.Printf("%d %d \n", v, sel_e[v])

		}
		for to := 0; to < n; to++ {
			if g[v][to] < min_e[to] {
				min_e[to] = g[v][to]
				sel_e[to] = v
			}
		}
	}
	return result
}
func main() {
	var n, m int
	var g [][]int
	var a, b, c int
	input.Scanf("%d%d", &n, &m)
	for i := 0; i < n; i++ {
		var temp []int
		for j := 0; j < n; j++ {
			temp = append(temp, 100000)
		}
		g = append(g, temp)
	}
	for i := 0; i < m; i++ {
		input.Scanf("%d %d %d", &a, &b, &c)
		g[a][b] = c
		g[b][a] = c

	}
	result := Prim(n, g)
	fmt.Printf("%d", result)
}

