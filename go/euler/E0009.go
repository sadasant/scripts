package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"math"
)

func solution(n int) interface{} {
	var a, b, c float64
	n64 := float64(n)
	for c = 3.0; c < n64; c++ {
		for b = c - 1; b > 1; b-- {
			for a = b - 1; a > 1; a-- {
				if a+b+c != n64 {
					continue
				}
				if math.Pow(a, 2.0)+math.Pow(b, 2.0) == math.Pow(c, 2.0) {
					return int(a * b * c)
				}
			}
		}
	}
	return nil
}

func main() {
	euler.Init(9, "There exists exactly one Pythagorean triplet for which a + b + c = 1000.")
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution, 1000)
}
