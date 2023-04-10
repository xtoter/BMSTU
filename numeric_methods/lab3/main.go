package main

import (
	"fmt"
	"math"
)

var n int = 10
var first int = 0
var end int = 1
var h float64
var matrix [][]float64

func fx(x float64) float64 {
	return math.Exp(x)
}
func fx1(x float64) float64 {
	return math.Sin(x) * math.Cos(x/2)
}
func newX(n int, a, b float64) ([]float64, []float64) {
	h := (b - a) / float64(n)
	var x, y []float64
	for i := 0; i < n+1; i++ {
		x = append(x, (a + float64(i)*h))
	}
	for _, val := range x {
		y = append(y, fx(val))
	}
	return x, y
}
func findD(y []float64, h float64, n int) []float64 {
	coef := 3 / (h * h)
	var d []float64
	for i := 1; i < n; i++ {
		d = append(d, coef*(y[i+1]-2*y[i]+y[i-1]))
	}
	return d
}
func findbaseX(x []float64, h float64) []float64 {
	var basex []float64
	for i := 1; i < len(x); i++ {
		basex = append(basex, float64(first)+(float64(i)-0.5)*h)
	}
	return basex
}
func find_a_b_d(c, y []float64, h float64, n int) ([]float64, []float64, []float64) {
	var a, b, d []float64
	for i := 1; i < n+1; i++ {
		a = append(a, y[i-1])
		val_b := (y[i]-y[i-1])/h - (h/3)*(c[i]+2*c[i-1])
		b = append(b, val_b)
		val_d := (c[i] - c[i-1]) / 3 * h
		d = append(d, val_d)
	}
	return a, b, d
}
func find_answer(A, B, C, D, basex, x []float64, n int) []float64 {
	var S []float64
	for i := 0; i < n; i++ {
		value := A[i] + C[i]*(basex[i]-x[i])*(basex[i]-x[i]) + B[i]*(basex[i]-x[i]) + D[i]*(basex[i]-x[i])*(basex[i]-x[i])*(basex[i]-x[i])
		S = append(S, value)
	}
	return S
}
func finddifference(basex, S []float64) ([]float64, []float64) {
	var mas, dif []float64
	for i := 0; i < len(basex); i++ {
		y := fx(basex[i])
		mas = append(mas, y)
		dif = append(dif, math.Abs(y-S[i]))
	}
	return dif, mas
}
func findbaseX2(x []float64, h float64) []float64 {
	var basex []float64
	for i := 1; i < len(x); i++ {
		basex = append(basex, float64(first)+float64(i)*h)
	}
	return basex
}
func createnewmatr(matrix [][]float64) ([]float64, []float64, []float64) {
	n := len(matrix)
	var a, b, c []float64
	b = append(b, matrix[0][0])
	c = append(c, matrix[0][1])
	num := 0
	for i := 1; i < n-1; i++ {
		num = 0
		for _, j := range matrix[i] {
			if j != 0 {
				if num == 0 {
					a = append(a, j)
					num += 1
				} else if num == 1 {
					b = append(b, j)
					num += 1
				} else {
					c = append(c, j)
					break
				}
			}
		}
	}
	a = append(a, matrix[n-1][n-2])
	b = append(b, matrix[n-1][n-1])
	return a, b, c
}
func straightstroke(a, b, c, d []float64, matrix [][]float64) ([]float64, []float64) {
	n := len(matrix)
	var alpha, beta []float64
	alpha0 := float64(-c[0] / b[0])
	beta0 := (d[0] / b[0])
	alpha = append(alpha, alpha0)
	beta = append(beta, beta0)
	for i := 1; i < len(matrix)-1; i++ {
		alpha = append(alpha, -c[i]/(a[i-1]*alpha[i-1]+b[i]))
		beta = append(beta, (d[i]-a[i-1]*beta[i-1])/(a[i-1]*alpha[i-1]+b[i]))
	}
	beta = append(beta, (d[n-1]-a[n-2]*beta[n-2])/(a[n-2]*alpha[n-2]+b[n-1]))
	return alpha, beta
}
func reversestroke(alpha, beta []float64, n int) []float64 {
	x := make([]float64, n)
	x[n-1] = beta[n-1]
	for i := n - 2; i > -1; i-- {
		x[i] = alpha[i]*x[i+1] + beta[i]
	}
	return x
}

/*

 */
func main() {
	h = (float64(end) - float64(first)) / float64(n)
	matrix = [][]float64{[]float64{4, 1, 0, 0, 0, 0, 0, 0, 0},
		[]float64{1, 4, 1, 0, 0, 0, 0, 0, 0},
		[]float64{0, 1, 4, 1, 0, 0, 0, 0, 0},
		[]float64{0, 0, 1, 4, 1, 0, 0, 0, 0},
		[]float64{0, 0, 0, 1, 4, 1, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 4, 1, 0, 0},
		[]float64{0, 0, 0, 0, 0, 1, 4, 1, 0},
		[]float64{0, 0, 0, 0, 0, 0, 1, 4, 1},
		[]float64{0, 0, 0, 0, 0, 0, 0, 1, 4}}
	a, b, c := createnewmatr(matrix)
	x, y := newX(n, float64(first), float64(end))
	d := findD(y, h, n)
	alpha, beta := straightstroke(a, b, c, d, matrix)
	C := reversestroke(alpha, beta, len(matrix))
	var temp []float64
	temp = append(temp, 0)
	temp = append(temp, C...)
	temp = append(temp, 0)
	C = temp
	basex := findbaseX(x, h)
	A, B, D := find_a_b_d(C, y, h, n)
	xs := findbaseX2(x, h)
	S1 := find_answer(A, B, C, D, xs, x, n)
	dif1, newY1 := finddifference(xs, S1)
	S := find_answer(A, B, C, D, basex, x, n)
	dif, newY := finddifference(basex, S)
	fmt.Println("X                       Y                 Spline                    Delta")
	for i := 0; i < len(xs); i++ {

		fmt.Printf("x: %3.15f y: %3.15f Spline: %3.15f D: %3.15f\n", basex[i], newY[i], S[i], dif[i])
		fmt.Printf("x: %3.15f y: %3.15f Spline: %3.15f D: %3.15f\n", xs[i], newY1[i], S1[i], dif1[i])
	}

	fmt.Println()
	//fmt.Println("X                       Y                 Spline                    Delta")
	//for i := 0; i < len(basex); i++ {
	//	fmt.Printf("x: %3.15f y: %3.15f Spline: %3.15f D: %3.15f\n", basex[i], newY[i], S[i], dif[i])
	//}
	//что-то произошло с узлами интерполяции
}
