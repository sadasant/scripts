// Daniel R. (sadasant.com)
// 21/09/2012
//
// Problem:
// http://projecteuler.net/problem=2
//
// How to run:
//   go run E0002.go 4000000
//

package main

import (
	"flag"
	"strconv"
)

func main() {
	flag.Parse()
	n, _ := strconv.Atoi(flag.Arg(0))

	a, b, sum := 0, 1, 0

	for b < n {
		b, a = a+b, b
		if b%2 == 0 {
			sum += b
		}
	}

	println(sum)
}
