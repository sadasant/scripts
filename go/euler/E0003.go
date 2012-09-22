// Daniel R. (sadasant.com)
// 03/02/2012
//
// Problem:
//   http://projecteuler.net/problem=3
//
// How to run:
//   go run E0003.go 600851475143
//

package main

import (
	"flag"
	"math"
	"strconv"
)

func main() {
	flag.Parse()
	n, _ := strconv.ParseInt(flag.Arg(0), 10, 64)

	nmax := int64(math.Floor(math.Sqrt(float64(n)))) / 3
	primes := make([]int64, nmax)
	factor, i := int64(0), int64(2)

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

	println(factor)
}
