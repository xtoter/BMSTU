package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/skorobogatov/input"
)

type graph struct {
	tops  []int
	ribs1 []int
	ribs2 []int
}

func ribssearch(cur graph) graph {
	for i := len(cur.tops) - 1; i >= 0; i-- {
		for v := i - 1; v >= 0; v-- {
			if cur.tops[i]%cur.tops[v] == 0 {
				temp := 1
				for u := i - 1; u > v; u-- {
					if cur.tops[i]%cur.tops[u] == 0 && cur.tops[u]%cur.tops[v] == 0 {
						temp = 0
						break
					}
				}
				if temp != 0 {
					cur.ribs1 = append(cur.ribs1, cur.tops[i])
					cur.ribs2 = append(cur.ribs2, cur.tops[v])
				}
			}
		}
	}
	return cur
}
func dividerstest(x int) []int {
	var a []int
	for i := 1; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			if i != x/i {
				a = append(a, i)
				a = append(a, x/i)
			} else {
				a = append(a, i)
			}
		}
	}
	return a
}
func dividers(x int) []int {
	var a []int
	for i := x; i > 0; i-- {
		if x%i == 0 {
			a = append(a, i)
		}
	}
	return a
}
func main() {
	var cur graph
	var n int
	input.Scanf("%d", &n)
	fmt.Println("graph {")
	cur.tops = dividerstest(n)
	sort.Ints(cur.tops)
	cur = ribssearch(cur)
	for i := len(cur.tops) - 1; i >= 0; i-- {
		fmt.Println(cur.tops[i])
	}
	for i := 0; i < len(cur.ribs1); i++ {
		fmt.Println(cur.ribs1[i], "--", cur.ribs2[i])
	}
	fmt.Println("}")
}

