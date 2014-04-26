package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"math/big"
)

func solution(n int) int {
	fact := big.NewInt(1).MulRange(1, int64(n))
	str := fact.String()
	sum := 0
	for _, v := range str {
		if i := int(v - '0'); i > 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	euler.Init(20, "Find the sum of the digits in the number 100!")
	euler.PrintTime("For 10  | Result: %v, Nanoseconds: %d\n", solution, 10)
	euler.PrintTime("For 100 | Result: %v, Nanoseconds: %d\n", solution, 100)
}
