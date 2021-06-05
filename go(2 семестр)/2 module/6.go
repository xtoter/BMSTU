package main

import (
	"fmt"
	"math"

	"github.com/skorobogatov/input"
)

type point struct {
	x int
	y int
}

func dist(a, b point) float64 {
	var z float64
	z = math.Sqrt(float64((b.x-a.x)*(b.x-a.x) + (b.y-a.y)*(b.y-a.y)))
	return z
}
func Prim(n int, g [][]float64) float64 {
	var result float64
	var used []bool
	var sel_e []int
	var min_e []float64
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
	var n int
	var g [][]float64
	var tempt []point
	var a, b int
	input.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var temp []float64
		for j := 0; j < n; j++ {
			temp = append(temp, 100000)
		}
		g = append(g, temp)
	}
	for i := 0; i < n; i++ {
		var z point
		input.Scanf("%d %d", &a, &b)
		z.x = a
		z.y = b
		tempt = append(tempt, z)

	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = dist(tempt[i], tempt[j])
		}
	}
	result := Prim(n, g)
	fmt.Printf("%.2f", result)
}

