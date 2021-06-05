package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

type Vertex struct {
	top    int
	weight int
}

func deikstra(s, n int, g []map[int]int) []int {

	//vector<int> d (n, INF),  p (n);
	var d []int
	var u []bool
	var p []int
	for i := 0; i < n; i++ {
		d = append(d, 10000)
		p = append(p, 0)
		u = append(u, false)
	}
	d[s] = 0
	//vector<char> u (n);

	for i := 0; i < n; i++ {
		v := -1
		for j := 0; j < n; j++ {
			if !u[j] && (v == -1 || d[j] < d[v]) {
				v = j
			}
		}
		if d[v] == 10000 {
			break
		}
		u[v] = true

		for j := 0; j < n; j++ {
			to := j
			var len int
			if g[v][j] == 0 {
				len = 10000
			} else {
				len = 1
			}

			if d[v]+len < d[to] {
				d[to] = d[v] + len
				p[to] = v
			}
		}
	}
	return d
}
func main() {
	var n, m, a, b, k int
	input.Scanf("%d%d", &n, &m)
	var g []map[int]int
	for i := 0; i < n; i++ {
		temp := make(map[int]int)

		g = append(g, temp)
	}
	for i := 0; i < m; i++ {
		input.Scanf("%d%d", &a, &b)
		g[a][b] = 1
		g[b][a] = 1

	}
	/*
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(g[i][j], " ")
			}
			fmt.Print("\n")
		}*/

	input.Scanf("%d", &k)
	var tops []int
	for i := 0; i < k; i++ {
		input.Scanf("%d", &a)
		tops = append(tops, a)
	}

	c := deikstra(tops[0], n, g)
	//fmt.Println(c)
	for j := 1; j < k; j++ {
		d := deikstra(tops[j], n, g)
		//fmt.Println(d)
		for i := 0; i < len(c); i++ {
			if c[i] != d[i] {
				c[i] = 10000
			}
		}
	}
	//fmt.Println(c)
	var r []int
	for i := 0; i < len(c); i++ {
		if c[i] != 10000 {
			r = append(r, i)
		}
	}
	if len(r) > 0 {
		for i := 0; i < len(r); i++ {
			fmt.Printf("%d ", r[i])
		}
	} else {
		fmt.Print("-")
	}
}

