package main

import (
	"fmt"
	"math"
)

func fx(x float64) float64 {
	return math.Sin(x)*math.Cos(x/2)
}
func fx1(x float64) float64 {
	return math.Sin(x) * math.Cos(x/2)
}
func middleRectangle(a, b float64, n int, funcX func(x float64) float64) float64 {
	h := (b - a) / float64(n)
	s := 0.0
	for i := 0; i < n; i++ {
		s += funcX(a + h*float64(i) + h*0.5)
	}
	return h * s
}
func trapezoid(a, b float64, n int, funcX func(x float64) float64) float64 {
	h := (b - a) / float64(n)
	s := 0.0
	for i := 1; i < n; i++ {
		s += funcX(a + h*float64(i))
	}
	return h * (((funcX(a) + funcX(b)) / 2.0) + s)
}
func simpson(a, b float64, n int, funcX func(x float64) float64) float64 {
	h := (b - a) / float64(n)
	s1, s2 := 0.0, 0.0
	for i := 0; i < n; i++ {
		s1 += funcX(a + h*float64(i) + h/2)
	}
	for i := 1; i < n; i++ {
		s2 += funcX(a + h*float64(i))
	}
	return h / 6 * (funcX(a) + funcX(b) + 4*s1 + 2*s2)
}
func richardson(h1, h2 float64, k int) float64 {
	return (h1 - h2) / (math.Pow(float64(2), float64(k)) - 1.0)
}
func calculate(eps float64, k int, f func(a, b float64, n int, x func(x float64) float64) float64, a, b float64, x func(x float64) float64) {
	n := 1
	iteration := 0
	h1, h2 := 0.0, 0.0
	r := 99999999.0
	for math.Abs(r) >= eps {
		n *= 2
		h2 = h1
		h1 = f(a, b, n, x)
		r = richardson(h1, h2, k)
		iteration++
	}
	fmt.Println("Iterations: ", iteration)
	fmt.Println("Result: ", h1)
	fmt.Println("Richardson: ", r)
	fmt.Println("with Richardson: ", h1+r)
}
func main() {
	for eps := 0.1; eps > 0.000001; eps /= 10 {
		fmt.Println("\033[34m", "eps: ", eps, "\033[0m")
		fmt.Println("\033[31m", "\ttrapezoid: ", "\033[0m")
		calculate(eps, 2, trapezoid, 0, 1, fx)
		fmt.Println("\033[31m", "\tmiddleRectangle: ", "\033[0m")
		calculate(eps, 2, middleRectangle, 0, 1, fx)
		fmt.Println("\033[31m", "\tSimpson: ", "\033[0m")
		calculate(eps, 4, simpson, 0, 1, fx)
		fmt.Println("\n")
	}
	///
	fmt.Println("\n---------------------------------------------------\n")
	for eps := 0.1; eps > 0.000001; eps /= 10 {
		fmt.Println("\033[34m", "eps: ", eps, "\033[0m")
		fmt.Println("\033[31m", "\ttrapezoid: ", "\033[0m")
		calculate(eps, 2, trapezoid, 0, math.Pi, fx1)
		fmt.Println("\033[31m", "\tmiddleRectangle: ", "\033[0m")
		calculate(eps, 2, middleRectangle, 0, math.Pi, fx1)
		fmt.Println("\033[31m", "\tSimpson: ", "\033[0m")
		calculate(eps, 4, simpson, 0, math.Pi, fx1)
		fmt.Println("\n")
	}
}
