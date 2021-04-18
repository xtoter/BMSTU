package main

import (
	"fmt"
	"math"

	"github.com/skorobogatov/input"
)

type data struct {
	timer   int
	tin     []int
	fup     []int
	used    []bool
	G       [][]int
	counter int
}

func dfs(v int, p int, x data) data {
	x.used[v] = true
	x.timer++
	x.fup[v] = x.timer
	x.tin[v] = x.fup[v]
	for i := 0; i < len(x.G[v]); i++ {
		cur := x.G[v][i]
		if cur != p {
			if x.used[cur] {
				x.fup[v] = int(math.Min(float64(x.fup[v]), float64(x.tin[cur])))
			} else {
				x = dfs(cur, v, x)
				x.fup[v] = int(math.Min(float64(x.fup[v]), float64(x.fup[cur])))
				if x.fup[cur] > x.tin[v] {

					x.counter++
				}
			}
		}
	}
	return x
}

func find_bridges(x data) data {
	x.timer = 0
	for i := 0; i < len(x.G); i++ {
		x.used[i] = false
	}
	for i := 0; i < len(x.G); i++ {
		if !x.used[i] {
			x = dfs(i, -1, x)
		}
	}
	return x
}
func initialization(x data, n int) data {
	x.G = make([][]int, n)
	x.tin = make([]int, len(x.G))
	x.fup = make([]int, len(x.G))
	x.used = make([]bool, len(x.G))
	x.counter = 0
	return x
}
func main() {
	var x data
	var a, b, n, m int
	input.Scanf("%d%d", &n, &m)
	x = initialization(x, n)
	for i := 0; i < m; i++ {
		input.Scanf("%d%d", &a, &b)
		x.G[a] = append(x.G[a], b)
		x.G[b] = append(x.G[b], a)
	}

	x = find_bridges(x)
	fmt.Println(x.counter)
}
