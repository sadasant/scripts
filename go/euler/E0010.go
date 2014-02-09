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

func main() {
	euler.Init(10, "Find the sum of all the primes below N.")
	euler.PrintTime("N = 10  | Result: %v, Nanoseconds: %d\n", solution, 10)
	euler.PrintTime("N = 2e6 | Result: %v, Nanoseconds: %d\n", solution, 2e6)
}
