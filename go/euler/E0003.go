package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"math"
)

func solution(n int) int64 {
	var m, l int64
	n64 := int64(n)
	m = 3
	for int64(math.Sqrt(float64(m))) < n64 {
		if n64%m == 0 {
			n64, l = n64/m, m
		}
		m += 2
	}
	return l
}

func main() {
	euler.Init(3, "What is the largest prime factor of the number 600851475143 ?")
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution, 600851475143)
}
