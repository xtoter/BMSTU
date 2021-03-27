package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

func compare(x int, y int) int {
	var n1, n2 int
	n1 = 10
	for x/n1 >= 1 {
		n1 = n1 * 10
	}
	n2 = 10
	for y/n2 >= 1 {
		n2 = n2 * 10
	}
	if (n1*y + x) > (n2*x + y) {
		return 1
	} else {
		return -1
	}
}
func main() {
	var n, tempright, right, templeft, left, kol int
	input.Scanf("%d", &n)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		input.Scanf("%d", &x[i])
	}
	right = n - 1
	kol = 0
	for (tempright*templeft > 0) || (kol == 0) {
		templeft = 0
		tempright = 0
		for i := left; i < right; i++ {
			if compare(x[i], x[i+1]) == 1 {
				x[i], x[i+1] = x[i+1], x[i]
				tempright = i
			}
		}
		right = tempright
		for i := right; i > left; i-- {
			if compare(x[i-1], x[i]) == 1 {
				x[i-1], x[i] = x[i], x[i-1]
				templeft = i
			}
		}
		left = templeft
		kol++
	}
	for i := 0; i < n; i++ {
		fmt.Print(x[i])
	}
}
