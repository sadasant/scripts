// Daniel R. (sadasant.com)
// 05/02/2012
//
// Problem:
//   http://projecteuler.net/problem=5
//
// How to run:
//   ./Euler_0005 20
//

package main

import (
	"fmt"
	"flag"
	"strconv"
)

// Euclidean Algorithm
func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {

	// Input
	flag.Parse()
	max, _ := strconv.Atoi(flag.Arg(0))

	// Variables
	mults := make([]int, max)
	lcm := 1

	// process:
	// gcd only for primes, but from last to first (?)
L:
	for i := max; i > 0; i-- {
		for j, v := range mults {
			if v == 0 {
				mults[j] = i
				lcm = (lcm * i) / gcd(lcm, i)
				break
			} else if v%i == 0 {
				break L
			}
		}
	}

	// Output
	fmt.Println(lcm)
}
