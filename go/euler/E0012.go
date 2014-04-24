package main

import "github.com/sadasant/scripts/go/euler/euler"

func solution(top int) interface{} {
	primes := euler.PrimesUpTo(top)

	var a, t, tt, exp, facts int
	a, t = 1, 1

	for facts < top {
		facts = 1
		a += 1
		t += a
		tt = t
		for _, v := range primes {
			if v*v > tt {
				facts *= 2
				break
			}
			exp = 1
			for tt%v == 0 {
				exp++
				tt = tt / v
			}
			if exp > 1 {
				facts *= exp
			}
			if tt == 1 {
				break
			}
		}
	}
	return t
}

func main() {
	euler.Init(12, "What is the value of the first triangle number to have over N divisors?")
	euler.PrintTime("N = 5   | Result: %v, Nanoseconds: %d\n", solution, 5)
	euler.PrintTime("N = 500 | Result: %v, Nanoseconds: %d\n", solution, 500)
}
