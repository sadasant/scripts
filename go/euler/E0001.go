package main

import "github.com/sadasant/scripts/go/euler/euler"

func solution(v ...int) interface{} {
	var s int
	for i := 0; i < v[0]; i++ {
		if i%3*i%5 == 0 {
			s += i
		}
	}
	return s
}

func main() {
	euler.Init(1, "Find the sum of all the multiples of 3 or 5 below 1000.")
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution, 1000)
}
