package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

func partition(less func(i, j int) bool, swap func(i, j int), low int, high int) int {
	i := low
	j := low
	for j < high {
		if less(j, high) {
			swap(i, j)
			i = i + 1
		}
		j = j + 1
	}
	swap(i, high)
	return i
}
func quicksortrec(less func(i, j int) bool, swap func(i, j int), low int, high int) {
	if low < high {
		q := partition(less, swap, low, high)
		quicksortrec(less, swap, low, q-1)
		quicksortrec(less, swap, q+1, high)
	}

}
func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	quicksortrec(less, swap, 0, n-1)
}
func main() {

	var n int
	input.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		input.Scanf("%d", &a[i])
	}
	swap := func(i int, j int) {
		a[i], a[j] = a[j], a[i]
	}
	less := func(i int, j int) bool {
		return a[i] < a[j]
	}
	qsort(n, less, swap)
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
}
