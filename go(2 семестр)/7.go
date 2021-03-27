package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

type frac struct {
	a int
	b int
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
func gcd(a, b int) int {
	a, b = abs(a), abs(b)
	for b > 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}
func normalize(f frac) frac {
	a, b := f.a, f.b
	if a < 0 && b < 0 {
		a, b = abs(a), abs(b)
	} else if a < 0 || b < 0 {
		a, b = -abs(a), abs(b)
	}
	k := gcd(a, b)
	a /= k
	b /= k
	return frac{a, b}
}
func print(example [][]frac, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n+1; j++ {
			fmt.Printf("%d / %d  ", example[i][j].a, example[i][j].b)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n next \n")
}
func subtract(example [][]frac, start, n int) (decision [][]frac) {
	//fmt.Printf("удалить")
	for i := start + 1; i < n; i++ {
		if example[i][start].a != 0 {
			for j := start; j < n+1; j++ {

				example[i][j].a = example[i][j].a*example[start][j].b - example[start][j].a*example[i][j].b
				example[i][j].b = example[i][j].b * example[start][j].b
			}
		}
	}
	return example
}
func sum(a, b frac) frac {
	var c frac

	c.b = a.b * b.b
	c.a = a.a*b.b + b.a*a.b

	c = normalize(c)

	return c
}
func replace(example [][]frac, i, n int) {
	if example[i][i].a == 0 {
		x := -1
		for j := i + 1; j < n; j++ {
			if x == -1 && example[j][i].a != 0 {
				x = j
			}
		}
		if x != -1 {
			for k := 0; k < n+1; k++ {
				temp := frac{example[i][k].a, example[i][k].b}
				example[i][k].a = example[x][k].a
				example[i][k].b = example[x][k].b
				example[x][k].a = temp.a
				example[x][k].b = temp.b
			}

		}
	}
}
func exception(example [][]frac, n int) (decision [][]frac) {
	for row, column := 0, 0; row < n; row++ {
		//print(example, n)
		if example[row][row].a == 0 {
			replace(example, row, n)
			//fmt.Printf("Замена")
			//print(example, n)
		}
		for i := row; i < n; i++ {
			//fmt.Printf("зашел\n")

			temp := frac{example[i][row].a, example[i][row].b}
			if temp.a != 0 {
				for j := column; j < n+1; j++ {

					example[i][j].a = example[i][j].a * temp.b
					example[i][j].b = example[i][j].b * temp.a
					example[i][j] = normalize(example[i][j])
				}
			}

		}
		//print(example, n)
		example = subtract(example, row, n)
		//print(example, n)
		column++
	}
	return example
}
func multiplication(a, b frac) frac {
	c := frac{1, 1}
	c.a = a.a * b.a
	c.b = a.b * b.b
	c = normalize(c)
	return c

}
func check(example [][]frac, n int) bool {
	x := true
	for i := 0; i < n; i++ {
		if x && example[i][i].a == 0 {
			x = false
		}
	}
	return x
}
func substitution(example [][]frac, n int) []frac {
	var answer []frac
	answer = append(answer, frac{1, 1})
	for i := n - 1; i >= 0; i = i - 1 {
		temp := frac{0, 1}

		for j := i + 1; j < n+1; j++ {
			//fmt.Printf("%d / %d  ", example[i][j].a, example[i][j].b)
			temp = sum(temp, multiplication(example[i][j], multiplication(frac{-1, 1}, answer[n-j])))
			//fmt.Printf("ymnoz %d %d %d %d\n", example[i][j].a, example[i][j].b, answer[n-j].a, answer[n-j].b)
		}
		//fmt.Printf("\n")
		answer = append(answer, temp)
	}
	return answer
}
func main() {
	var n, temp int
	input.Scanf("%d", &n)
	var example [][]frac
	for i := 0; i < n; i++ {
		example = append(example, []frac{})
		for j := 0; j < n+1; j++ {
			input.Scanf("%d", &temp)
			example[i] = append(example[i], frac{temp, 1})
		}
	}
	example = exception(example, n)
	if check(example, n) {

		result := substitution(example, n)
		for i := n; i > 0; i-- {
			fmt.Printf("%d/%d\n", -1*result[i].a, result[i].b)
		}
	} else {
		fmt.Printf("No solution")
	}
}
