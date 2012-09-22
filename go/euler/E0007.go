// Daniel R. (sadasant.com)
// 11/02/2012
//
// Problem:
//   http://projecteuler.net/problem=7
//
// How to run:
//   go run E0007.go 10001
//

package main

import (
	"flag"
	"strconv"
)

func main() {

	// Input
	flag.Parse()
	n, _ := strconv.Atoi(flag.Arg(0))

	// Variables
	primes := make([]int64, n)
	var i int64 = 2

	// Process
	for l := 0; l < n; i++ {
		for j, v := range primes {
			if v == 0 {
				primes[j] = i
				l++
				break
			}
			if i%v == 0 {
				break
			}
		}
	}

	// Output
	println(primes[n-1])
}
