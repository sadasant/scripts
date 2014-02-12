package main

import "github.com/sadasant/scripts/go/euler/euler"

func Collatz(n int) int {
	var chain int
	for {
		chain++
		switch {
		case n == 1:
			return chain
		case n%2 == 0:
			n = n / 2
		default:
			n = 3*n + 1
		}
	}
}

func solution(v ...int) interface{} {
	var longest, longest_n, chain int
	for n := 1; n < v[0]; n++ {
		chain = Collatz(n)
		if chain > longest {
			longest = chain
			longest_n = n
		}
	}
	return longest_n
}

func main() {
	euler.Init(14, "Using the rules:\n\n"+
		"  n -> n/2 (n is even)\n"+
		"  n -> 3n + 1 (n is odd)\n\n"+
		"Which starting number, under one million, produces the longest chain?")
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution, 1e6)
}
