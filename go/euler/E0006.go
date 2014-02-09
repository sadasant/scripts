package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"math"
)

func solution(v ...int) interface{} {
	var n, sum_sq, sq_sum float64
	n = float64(v[0])
	for i := 1.0; i <= n; i++ {
		sum_sq += math.Pow(i, 2.0)
		sq_sum += i
	}
	return math.Pow(sq_sum, 2.0) - sum_sq
}

func main() {
	euler.Init(6, "Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.")
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution, 100)
}
