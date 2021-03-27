package main

import (
	"fmt"
	"math/big"

	"github.com/skorobogatov/input"
)

func binsearch(a1, a2, a3, a4 *big.Int, n int) (test1, test2, test3, test4 *big.Int) {
	if n == 0 {
		return big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1)

	}
	if n%2 == 1 {
		b1, b2, b3, b4 := binsearch(a1, a2, a3, a4, n-1)
		return new(big.Int).Add(new(big.Int).Mul(b1, a1), new(big.Int).Mul(b2, a3)),
			new(big.Int).Add(new(big.Int).Mul(b1, a2), new(big.Int).Mul(b2, a4)),
			new(big.Int).Add(new(big.Int).Mul(b3, a1), new(big.Int).Mul(b4, a3)),
			new(big.Int).Add(new(big.Int).Mul(b3, a2), new(big.Int).Mul(b4, a4))

	}
	b1, b2, b3, b4 := binsearch(a1, a2, a3, a4, n/2)
	return new(big.Int).Add(new(big.Int).Mul(b1, b1), new(big.Int).Mul(b2, b3)),
		new(big.Int).Add(new(big.Int).Mul(b1, b2), new(big.Int).Mul(b2, b4)),
		new(big.Int).Add(new(big.Int).Mul(b3, b1), new(big.Int).Mul(b4, b3)),
		new(big.Int).Add(new(big.Int).Mul(b3, b2), new(big.Int).Mul(b4, b4))

}
func search(n int) (test *big.Int) {
	temp1, _, _, _ := binsearch(big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(0), n-1)
	return temp1
}
func main() {
	var n int
	input.Scanf("%d", &n)
	fmt.Printf("%d", search(n))
}
