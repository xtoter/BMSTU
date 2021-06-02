package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

type data struct {
	a       []int
	b       []rune
	n, m, q int
}

func print(data data) {
	c := '"'
	fmt.Printf("digraph { \n rankdir = LR \n dummy [label = %c%c, shape = none] \n dummy -> %d\n", c, c, data.q)
	for i := 0; i < data.n; i++ {
		fmt.Printf("%d [shape = circle]\n", i)
	}
	for i := 0; i < data.n; i++ {
		for j := 0; j < data.m; j++ {
			fmt.Printf("%d -> %d  [label = %c%c(%c)%c]\n", i, data.a[i*data.m+j], c, 97+j, data.b[i*data.m+j], c)
		}
	}
	fmt.Printf("}")

}
func main() {
	var data data
	var temp1 int
	var temp2 rune
	input.Scanf("%d%d%d", &data.n, &data.m, &data.q)
	for i := 0; i < data.n*data.m; i++ {
		input.Scanf("%d", &temp1)
		data.a = append(data.a, temp1)
	}
	i := 0
	for i < data.n*data.m {
		input.Scanf("%c", &temp2)
		if temp2 != ' ' && temp2 != '\n' {

			data.b = append(data.b, temp2)
			i++
		}
	}
	print(data)
	//fmt.Print(a, b)
}

