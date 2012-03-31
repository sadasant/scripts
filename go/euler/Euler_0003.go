// Daniel R. (sadasant.com)
// 03/02/2012
//
// Problem:
//   http://projecteuler.net/problem=3
//
// How to run:
//   ./Euler_0003 600851475143
//

package main

import (
	"fmt"
	"flag"
	"strconv"
	"math"
)

func main() {

	// Input
	flag.Parse()
	n, _ := strconv.Atoi64(flag.Arg(0))

	// Variables
	nmax := int64(math.Floor(math.Sqrt(float64(n)))) / 3
	primes := make([]int64, nmax)
	factor, i := int64(0), int64(2)

	// Process
	for ; i < nmax; i++ {
		for j, v := range primes {
			if v == 0 {
				primes[j] = i
				if n%i == 0 {
					factor = i
				}
				break
			}
			if i%v == 0 {
				break
			}
		}
	}

	// Output
	fmt.Println(factor)
}
