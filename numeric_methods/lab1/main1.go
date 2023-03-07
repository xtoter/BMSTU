package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func mul_m(A [][]float32, x []float32) []float32 {
	var out []float32
	for i := 0; i < len(A); i++ {
		out = append(out, 0)
		for j := 0; j < len(A[i]); j++ {
			out[i] += A[i][j] * x[j]
		}
	}
	return out
}
func Tridiagonal_matrix_algorithm(A [][]float32, d []float32) []float32 {
	var alpha, beta, x []float32
	for i := 0; i < len(A); i++ {
		if i == 0 {
			alpha = append(alpha, -1*A[i][i+1]/A[i][i])
			beta = append(beta, d[i]/A[i][i])
		} else if i == len(A)-1 {
			alpha = append(alpha, 0)
			beta = append(beta, (d[i]-A[i][i-1]*beta[i-1])/(A[i][i]+alpha[i-1]*A[i][i-1]))
		} else {
			alpha = append(alpha, -1*A[i][i+1]/(A[i][i]+alpha[i-1]*A[i][i-1]))
			beta = append(beta, (d[i]-A[i][i-1]*beta[i-1])/(A[i][i]+alpha[i-1]*A[i][i-1]))
		}
	}
	//fmt.Println(alpha, beta)
	for i := 0; i < len(A); i++ {
		if i == 0 {
			x = append(x, float32(math.Abs(float64(beta[len(A)-1-i]))))
		} else {

			x = append(x, float32(math.Abs(float64(alpha[len(A)-1-i]*x[i-1]+beta[len(A)-1-i]))))
		}
	}
	var reverse []float32

	for i := len(x) - 1; i >= 0; i-- {
		reverse = append(reverse, x[i])
	}
	return reverse
}
func get_delta(a []float32, b []float32) []float32 {
	var out []float32
	for i := 0; i < len(a); i++ {
		out = append(out, float32(math.Abs(float64(a[i]-b[i]))))
	}
	return out
}
func main() {
	sizeA := 0
	var A [][]float32
	var d, wait_x []float32
	fmt.Scan(&sizeA)

	for i := 0; i < sizeA; i++ {
		var temp []float32
		for j := 0; j < sizeA; j++ {
			var temp_num float32
			fmt.Scan(&temp_num)
			temp = append(temp, temp_num)
		}
		A = append(A, temp)
	}
	for j := 0; j < sizeA; j++ {
		var temp_num float32
		fmt.Scan(&temp_num)
		d = append(d, temp_num)
	}
	for j := 0; j < sizeA; j++ {
		var temp_num float32
		fmt.Scan(&temp_num)
		wait_x = append(wait_x, temp_num)
	}
	//fmt.Println(d)
	_x := Tridiagonal_matrix_algorithm(A, d)
	fmt.Println("_x ", _x)
	//fmt.Println(mul_m(A, _x))
	delt_d := get_delta(mul_m(A, _x), d)
	fmt.Println("delta_d ", delt_d)

	// x := get_delta(_x, eps)
	// fmt.Println(x)
	var data []float32
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			data = append(data, A[i][j])
		}
	}
	tempdata := make([]float64, len(data))
	for i, v := range data {
		tempdata[i] = float64(v)
	}
	a := mat.NewDense(len(A), len(A), tempdata)
	var aInv mat.Dense
	err := aInv.Inverse(a)
	if err != nil {
		fmt.Println(err, aInv)
	}
	newdata := aInv.RawMatrix().Data
	var newA [][]float32
	for i := 0; i < sizeA; i++ {
		var temp []float32
		for j := 0; j < sizeA; j++ {
			temp = append(temp, float32(newdata[i*sizeA+j]))
		}
		newA = append(newA, temp)
	}
	//fmt.Println("a 1 ", newA, delt_d)
	delta_x := mul_m(newA, delt_d)
	//wait_x := []float32{1.0, 1.0, 1.0, 1.0}
	//wait_x := []float32{1.0, 2.0, 3.0, 4.0}
	eps_x := get_delta(wait_x, _x)
	fmt.Println("delta_x 1 ", delta_x)
	fmt.Println("delta_x 2", eps_x)
	//fmt.Println(A, d)
}
