package main

import (
	. "fmt"
	. "math"
)

func main() {
	var n, sum_sq, sq_sum float64

	Print("The difference between the sum of the squares and the square of the sum of the natural numbers up to ")
	Scan(&n)

	for i := 1.0; i <= n; i++ {
		sum_sq += Pow(i, 2.0)
		sq_sum += i
	}
	sq_sum = Pow(sq_sum, 2.0)

	Println("is", int(sq_sum - sum_sq))
}
