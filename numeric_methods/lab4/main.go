package main

import (
	"fmt"
	"math"
)

type callable func(float64) float64

func individualFunc(x float64) float64 {
	return math.Sin(x)*math.Cos(x/2)
}

func solveTridiagonalMatrix(a, b, c, d []float64) []float64 {
	alpha := []float64{-c[0] / b[0]}
	beta := []float64{d[0] / b[0]}
	n := len(d)
	x := make([]float64, n)

	for i := 1; i < n-1; i++ {
		alpha = append(alpha, -c[i]/(b[i]+a[1]*alpha[i-1]))
		beta = append(beta, (d[i]-a[i-1]*beta[i-1])/(b[i]+a[i-1]*alpha[i-1]))
	}

	beta = append(beta, (d[n-1]-a[n-2]*beta[n-2])/(b[n-1]+a[n-2]*alpha[n-2]))

	x[n-1] = (d[n-1] - a[n-2]*beta[n-2]) / (a[n-2]*alpha[n-2] + b[n-1])

	for i := n - 1; i > 0; i-- {
		x[i-1] = alpha[i-1]*x[i] + beta[i-1]
	}

	return x
}
func out(a [][]float64){
	fmt.Println(" i x	y		       y'		 eps")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%2.0d %3.2f %3.15f %3.15f %3.15f\n",i,a[i][0], a[i][1], a[i][2], a[i][3])
	}
}
func method1(f func(float64) float64, p float64, q float64, left float64, right float64, y0 float64, y1 float64, n int)[][]float64 {
	x := make([]float64, n+1)
	h := 1.0 / float64(n)
	for i := 0; i < n+1; i++ {
		x[i] = float64(i) * h
	}
	y := make([]float64, n+1)
	y[0] = y0
	y[n] = y1
	
	a := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		a[i] = 1 - p*h/2.0
	}
	b := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		b[i] = math.Pow(h, 2)*q - 2
	}
	c := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		c[i] = 1 + p*h/2.0
	}

	d := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		d[i] = math.Pow(h, 2) * math.Exp(x[i])
	}
	d[0] -= y0 * a[0]
	d[n-2] -= y1 * c[0]

	solution := solveTridiagonalMatrix(a, b, c, d)
	
	for i := 1; i < n; i++ {
		y[i] = solution[i-1]
	}
	eps := make([]float64, n+1)
	for i := 0; i < n+1; i++ {
		eps[i] = math.Abs(y[i] - f(x[i]))
	}

	data := make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		data[i] = make([]float64, 4)
		data[i][0] = x[i]
		data[i][1] = f(x[i])
		data[i][2] = y[i]
		data[i][3] = eps[i]
	}
	return data
	//fmt.Println(data)
}

func method2(f callable, p int, q int, left int, right int, y0 float64, y1 float64, n int) [][]float64 {
	h := float64(right-left) / float64(n)
	x := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		x[i] = float64(left) + float64(i)*h
	}
	y := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		y[i] = f(x[i])
	}

	y_0 := []float64{y0, y0 + h}
	y_1 := []float64{0, h}

	for i := 1; i <= n; i++ {
		y_0 = append(y_0, (y[i]*math.Pow(h, 2)+(2-float64(q)*math.Pow(h, 2))*y_0[i]-(1-float64(p)*h/2)*y_0[i-1])/(1+float64(p)*h/2))
	}

	for i := 1; i < n; i++ {
		y_1 = append(y_1, (2-float64(q)*math.Pow(h, 2))*y_1[i]-(1-float64(p)*h/2)*y_1[i-1]/(1+float64(p)*h/2))
	}

	c1 := (y1 - y_0[n]) / y_1[n]

	y_new := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		y_new[i] = y_0[i] + c1*y_1[i]
	}

	eps := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		eps[i] = math.Abs(y_new[i] - y[i])
	}
	data := make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		data[i] = make([]float64, 4)
		data[i][0] = x[i]
		data[i][1] =  y[i]
		data[i][2] = y_new[i]
		data[i][3] = eps[i]
	}
	return data
}

func main() {
	n := 10
	y0 := 3.0
	y1 := 2.5*math.E + 1/math.E
	p := 0
	q := -1

	// решение краевой задачи методом прогонки
	out(method1(individualFunc, float64(p), float64(q), 0, 1, y0, y1, n))
	// решение краевой задачи методом стрельбы
	out(method2(individualFunc, p, q, 0, 1, y0, y1, n))
}
