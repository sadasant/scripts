package main

import "github.com/sadasant/scripts/go/euler/euler"

func solution(v ...int) interface{} {
	n := v[0]
	P := make([]bool, n)
	s := 2
	for i := 3; i < n; i += 2 {
		if !P[i] {
			s += i
			for j := i; j < n; j += i {
				P[j] = true
			}
		}
	}
	return s
}

func solution2(v ...int) interface{} {
	var sum int
	for _, v := range euler.PrimesUpTo(v[0]) {
		sum += v
	}
	return sum
}

func main() {
	euler.Init(10, "Find the sum of all the primes below N.")
	euler.PrintTime("N = 10  | Using local Sieve without []int. Result: %v, Nanoseconds: %d\n", solution, 10)
	euler.PrintTime("N = 10  | Using packaged PrimesUpTo.       Result: %v, Nanoseconds: %d\n", solution2, 10)
	euler.PrintTime("N = 2e6 | Using local Sieve without []int. Result: %v, Nanoseconds: %d\n", solution, 2e6)
	euler.PrintTime("N = 2e6 | Using packaged PrimesUpTo.       Result: %v, Nanoseconds: %d\n", solution2, 2e6)
}
