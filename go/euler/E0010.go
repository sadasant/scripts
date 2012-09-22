// Daniel R. (sadasant.com)
// 12/02/2012
//
// Problem:
//   http://projecteuler.net/problem=10
//
// How to run:
//   go run E0010.go 2000000
//

package main

import (
	"flag"
	"strconv"
	M "math"
)

func main() {

	// Input
	flag.Parse()
	n, _ := strconv.ParseInt(flag.Arg(0), 10, 64)
	sq := int64(M.Sqrt(float64(n)))

	// Variables
	P := make([]bool, n)
	s := int64(0)

	// Process
	for i := int64(2); i < n; i++ {
		if !P[i] {
			P[i] = true
			s += i
			if i < sq {
				for j := i+i; j < n; j+=i {
					P[j] = true
				}
			}
		}
	}

	// Output
	println(s)
}
