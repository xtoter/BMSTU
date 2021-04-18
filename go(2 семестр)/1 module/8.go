package main

import (
	"fmt"
)

func main() {
	var example string
	var stack []int
	var count int
	fmt.Scanf("%s", &example)
	fmt.Print(example)
	hash := make(map[string]bool)
	exampleint := []rune(example)
	for i, symbol := range exampleint {
		if symbol == '(' {
			stack = append(stack, i)
		} else if symbol == ')' {
			temp := example[stack[len(stack)-1] : i+1]
			stack = stack[:len(stack)-1]
			if hash[temp] == false {
				hash[temp] = true
				count++
			}
		}
	}
	fmt.Print(count)
}
