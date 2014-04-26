package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"math"
	"strconv"
)

func solution(n float64) int {
	ns := strconv.FormatFloat(n, 'f', 12, 64)
	s := 0
	for _, v := range ns {
		i := int(v - '0')
		if i > 0 {
			s += i
		}
	}
	return s
}

func main() {
	euler.Init(16, "What is the sum of the digits of the number 2^1000?")
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution, math.Pow(2, 1000))
}
