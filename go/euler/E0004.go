package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"math"
	"strconv"
)

func reverse(s string) string {
	l := len(s)
	r := make([]rune, l)
	for _, v := range s {
		l--
		r[l] = v
	}
	return string(r[l:])
}

func solution(n int) interface{} {
	max := int(math.Pow10(n) - 1)
	min := 9 * max / 10
	for i := max; ; i-- {
		for j := max; j > min; j-- {
			p := strconv.Itoa(i * j)
			if p == reverse(p) {
				return p
			}
		}
	}
}

func main() {
	euler.Init(4, "Find the largest palindrome made from the product of two N-digit numbers.")
	euler.PrintTime("N = 2 | Result: %v, Nanoseconds: %d\n", solution, 2)
	euler.PrintTime("N = 3 | Result: %v, Nanoseconds: %d\n", solution, 3)
}
