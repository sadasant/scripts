// Daniel R. (sadasant.com)
// 22/09/2012
//
// Problem:
//   http://projecteuler.net/problem=12
//
// How to run:
//   go run E0012.go 500
//

package main

import (
	"flag"
	"strconv"
)

// Sieve of Erastosthenes
func primesTo(n int64) []int64 {
	bools := make([]bool, n+1)
	primes := make([]int64, n+1)
	primes[0] = 1
	c := 1
	for i := int64(2); i <= n; i++ {
		if !bools[i] {
			bools[i] = true
			primes[c] = i
			c++
			for j := i + i; j <= n; j += i {
				bools[j] = true
			}
		}
	}
	return primes[:c]
}

func main() {
	// Input
	flag.Parse()
	max, _ := strconv.ParseInt(flag.Arg(0), 10, 64)

	primes := primesTo(max)
	lenprimes := len(primes)

	var a, t, tt, exp, facts int64
	a, t, tt, exp = 1, 1, 0, 0

	for facts < max {
		facts = 1
		a = a + 1
		t = t + a // triangle
		tt = t
		for i := 1; i < lenprimes; i++ {
			if primes[i]*primes[i] > tt {
				facts *= 2
				break
			}
			exp = 1
			for tt%primes[i] == 0 {
				exp++
				tt = tt / primes[i]
			}
			if exp > 1 {
				facts *= exp
			}
			if tt == 1 {
				break
			}
		}
	}
	print(t)
}
