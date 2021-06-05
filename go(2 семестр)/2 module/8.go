package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

var time int
var count int
var stack []int

type graph struct {
	a    []int
	comp int
	T1   int
	low  int
}

func was(data []int, x int) bool {
	z := 0
	for i := 0; i < len(data); i++ {
		if data[i] == x {
			z++
			return false
		}
	}
	return true
}
func Tarjan(u []graph) []graph {
	for i := 0; i < len(u); i++ {
		u[i].T1 = 0
		u[i].comp = 0
	}
	for i := 0; i < len(u); i++ {
		if u[i].T1 == 0 {
			u = VisitVertex_Tarjan(u, i)
		}
	}
	return u
}
func VisitVertex_Tarjan(g []graph, v int) []graph {
	g[v].T1 = time
	g[v].low = time
	time++
	stack = append(stack, v)

	for j := 0; j < len(g[v].a); j++ {
		temp := g[v].a[j]
		if g[temp].T1 == 0 {
			g = VisitVertex_Tarjan(g, temp)
		}
		if g[temp].comp == 0 && g[v].low > g[temp].low {
			g[v].low = g[temp].low
		}
	}
	if g[v].T1 == g[v].low {
		for {
			u := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			g[u].comp = count
			if u == v {
				break
			}
		}
		count++
	}
	return g
}

func main() {
	time, count = 1, 1
	var n, m int
	input.Scanf("%d%d", &n, &m)
	var a []graph
	for i := 0; i < n; i++ {
		var temp graph
		a = append(a, temp)
	}

	for i := 0; i < m; i++ {
		var temp1, temp2 int
		input.Scanf("%d%d", &temp1, &temp2)
		a[temp1].a = append(a[temp1].a, temp2)

	}
	a = Tarjan(a)
	mark := make([][]byte, count)
	for i := 0; i < count; i++ {
		mark[i] = make([]byte, count)
	}
	component := make([][]int, count)
	for i := 0; i < n; i++ {
		component[a[i].comp] = append(component[a[i].comp], i)
	}

	var ta, tb []int
	for i := 0; i < n; i++ {
		for j := 0; j < len(a[i].a); j++ {
			temp := a[i].a[j]
			if (a[i].comp != a[temp].comp) && (mark[a[i].comp-1][a[temp].comp-1] != 1) {
				mark[a[i].comp-1][a[temp].comp-1] = 1
				ta = append(ta, a[i].comp)
				tb = append(tb, a[temp].comp)
			}
		}
	}
	var t []int
	for i := 1; i < count; i++ {
		t = append(t, i)
	}
	var result []int
	result = append(result, -1)
	//fmt.Print(ta, tb)
	for i := 0; i < len(t); i++ {
		if was(tb, t[i]) && (t[i] != result[len(result)-1]) {
			result = append(result, t[i])
		}
	}
	//fmt.Print(result)
	result = result[1:]

	var print []int
	for i := 0; i < len(result); i++ {
		print = append(print, component[result[i]][0])
	}
	for i := 0; i < len(print); i++ {
		fmt.Print(print[i], " ")
	}
	if len(print) == 0 {
		fmt.Print(0)
	}

}

