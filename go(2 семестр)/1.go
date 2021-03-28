package main

import (
	"fmt"
	"math"

	"github.com/skorobogatov/input"
)

func help(x uint64) (uint64, uint64) {
	for n, sum := uint64(0), uint64(0); ; n++ {
		k := 9 * uint64(math.Pow(float64(10), float64(n))) * (n + 1)
		sum = sum + k
		if sum > x {
			x = x - sum + k
			n++
			return n, x
		}
	}
}
func main() {
	var x uint64
	input.Scanf("%d", &n)
	a, b := help(n)
	number := (b / a) + uint64(math.Pow(float64(10), float64(a-1)))
	b = a - b%a
	for i := uint64(0); i < b-1; i++ {
		number = number / 10
	}
	result := (int)(number % 10)
	fmt.Printf("%d\n", result)
}
