package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"math"
)

func solution(v ...int) interface{} {
	var n, m, l int64
	n = int64(v[0])
	m = 3
	for int64(math.Sqrt(float64(m))) < n {
		if n%m == 0 {
			n, l = n/m, m
		}
		m += 2
	}
	return l
}

func main() {
	euler.Init(3, "What is the largest prime factor of the number 600851475143 ?")
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution, 600851475143)
}
