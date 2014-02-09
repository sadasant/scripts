package main

import "github.com/sadasant/scripts/go/euler/euler"

// Euclidean Algorithm
func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(v ...int) int {
	lcm := 1
	for _, v := range v {
		lcm = (lcm * v) / gcd(lcm, v)
	}
	return lcm
}

func main() {
	euler.Init(5, "What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?")
	euler.PrintTime("Index by index.     Result: %v, Nanoseconds: %d\n", lcm, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	euler.PrintTime("Generated sequence. Result: %v, Nanoseconds: %d\n", lcm, euler.Sequence(1, 20)...)
}
