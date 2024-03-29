package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

type automat struct {
	n, m, q      int
	delta, delta_new [][]int
	fi, fi_new [][]string
}

func inisialization(n, m, q int, delta [][]int, fi [][]string) automat {
	var cur automat
	cur.n = n
	cur.m = m
	cur.q = q
	cur.delta = delta
	cur.fi = fi
	var a [][]int
	for i := 0; i < n; i++ {
		var temp []int
		for j := 0; j < m; j++ {
			temp = append(temp, 0)
		}
		a = append(a, temp)
	}
	var b [][]string
	for i := 0; i < n; i++ {
		var temp []string
		for j := 0; j < m; j++ {
			temp = append(temp, "0")
		}
		b = append(b, temp)
	}
	cur.delta_new = a
	cur.fi_new = b
	return cur
}

func Find(a []int, x int) int {
	if a[x] == x {
		return x
	}
	a[x] = Find(a, a[x])
	return a[x]
}
func Union(a []int, x, y int) {
	x_new := Find(a, x)
	y_new := Find(a, y)
	if x_new == y_new {
		return
	}
	x_new, y_new = y_new, x_new
	a[x_new] = y_new
}
func Split(m int, pi []int, automat automat) int {
	m = automat.n
	var a []int
	for i := 0; i < m; i++ {
		a = append(a, i)
	}
	for i := 0; i < automat.n; i++ {
		for j := i + 1; j < automat.n; j++ {
			if pi[i] == pi[j] && Find(a, i) != Find(a, j) {
				eq := true
				for k := 0; k < automat.m; k++ {
					w1 := automat.delta[i][k]
					w2 := automat.delta[j][k]
					if pi[w1] != pi[w2] {
						eq = false
						break
					}
				}
				if eq {
					Union(a, i, j)
					m--
				}
			}
		}
	}
	for i := 0; i < automat.n; i++ {
		pi[i] = Find(a, i)
	}
	return m
}
func Split1(m int, pi []int, automat automat) int {
	m = automat.n
	var a []int
	for i := 0; i < m; i++ {
		a = append(a, i)
	}
	for i := 0; i < automat.n; i++ {
		for j := i + 1; j < automat.n; j++ {
			if Find(a, i) != Find(a, j) {
				eq := true
				for k := 0; k < automat.m; k++ {
					if automat.fi[i][k] != automat.fi[j][k] {
						eq = false
						break
					}
				}
				if eq {
					Union(a, i, j)
					m--
				}
			}
		}
	}
	for i := 0; i < automat.n; i++ {
		pi[i] = Find(a, i)
	}
	return m
}
func AuftenkampHohn(automat automat) automat {
	var pi []int
	for i := 0; i < automat.n; i++ {
		pi = append(pi, 0)
	}
	m1, m2 := -1, -1
	m1 = Split1(m1, pi, automat)
	for {
		m1 = Split(m1, pi, automat)
		if m1 == m2 {
			break
		}
		m2 = m1

	}
	var pi1, pi2 []int
	for i := 0; i < automat.n; i++ {
		pi1 = append(pi1, 0)
		pi2 = append(pi2, 0)
	}
	a := 0
	for i := 0; i < automat.n; i++ {
		if pi[i] == i {
			pi2[a] = i
			pi1[i] = a
			a++
		}
	}
	automat.n = m1
	automat.q = pi1[pi[automat.q]]
	var p [][]string
	for i := 0; i < automat.n; i++ {
		var temp []string
		for j := 0; j < automat.m; j++ {
			temp = append(temp, "0")
		}
		p = append(p, temp)
	}
	temp := automat.fi
	for i := 0; i < automat.n; i++ {
		for j := 0; j < automat.m; j++ {
			automat.delta[i][j] = pi1[pi[automat.delta[pi2[i]][j]]]
			p[i][j] = temp[pi2[i]][j]
		}

	}
	automat.fi = p
	return automat
}
func Print(automat automat) {
	automat = prepare(automat)
	fmt.Printf("digraph {\nrankdir = LR\ndummy [label = \"\", shape = none]\n")
	for i := 0; i < automat.n; i++ {
		fmt.Printf("%d  [shape = circle]\n", i)
	}
	fmt.Printf("dummy ->  %d\n", automat.q)
	for i := 0; i < automat.n; i++ {
		for j := 0; j < automat.m; j++ {
			fmt.Printf("%d -> %d [label = \"%c(%s)\"]\n", i, automat.delta[i][j], 97+j, automat.fi[i][j])
		}

	}
	fmt.Printf("}")
}
func dfs(begin int, automat automat) (int, []int) {
	var used []int
	for i := 0; i < automat.n; i++ {
		used = append(used, -1)
	}
	index := VisitVertex(used, begin, 0, automat)
	return index, used
}
func VisitVertex(used []int, begin int, index int, automat automat) int {
	used[begin] = index
	index++
	for i := 0; i < automat.m; i++ {
		if used[automat.delta[begin][i]] == -1 {
			index = VisitVertex(used, automat.delta[begin][i], index, automat)
		}
	}
	return index
}
func prepare(automat automat) automat {

	index, used := dfs(automat.q, automat)
	for i := 0; i < automat.n; i++ {
		if used[i] != -1 {
			automat.fi_new[used[i]] = automat.fi[i]
			for j := 0; j < automat.m; j++ {
				automat.delta_new[used[i]][j] = used[automat.delta[i][j]]
			}
		}
	}
	automat.q = 0
	automat.n = index
	automat.delta = automat.delta_new
	automat.fi = automat.fi_new
	return automat
}
func main() {
	var n, m, q int
	input.Scanf("%d%d%d", &n, &m, &q)
	var delta [][]int
	for i := 0; i < n; i++ {
		var temp []int
		for j := 0; j < m; j++ {
			var temmp int
			input.Scanf("%d", &temmp)
			temp = append(temp, temmp)
		}
		delta = append(delta, temp)
	}
	var fi [][]string
	for i := 0; i < n; i++ {
		var temp []string
		for j := 0; j < m; j++ {
			var temmp string
			input.Scanf("%s", &temmp)
			temp = append(temp, temmp)
		}
		fi = append(fi, temp)
	}
	automat := inisialization(n, m, q, delta, fi)
	automat = AuftenkampHohn(automat)
	Print(automat)
}

