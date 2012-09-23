// Daniel R. (sadasant.com)
// 22/09/2012
//
// Problem:
//   http://projecteuler.net/problem=9
//
// How to run:
//   go run E0009.go 1000
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
	n, _ := strconv.ParseFloat(flag.Arg(0), 64)
	a, b, c, t := 0.0, 0.0, 0.0, false

	// Process
L:	for c = 3.0; c < n; c++ {
		for b = c-1; b > 1; b-- {
			for a = b-1; a > 1; a-- {
				if a + b + c != n { continue }
				t = M.Pow(a, 2.0) + M.Pow(b, 2.0) == M.Pow(c, 2.0)
				if t { break L }
			}
		}
	}

	// Output
	if t {
		println(a,b,c,int(a*b*c))
	}
}
