package main

import "github.com/sadasant/scripts/go/euler/euler"

func solution(n int) int {
	a, b, s := 0, 1, 0
	for b < n {
		b, a = a+b, b
		if b%2 == 0 {
			s += b
		}
	}
	return s
}

func main() {
	euler.Init(2, "By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.")
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution, int(4e6))
}
