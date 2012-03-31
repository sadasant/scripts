// Daniel R. (sadasant.com)
// 01/02/2012
//
// Problem:
// http://projecteuler.net/problem=2
//
// How to run:
//   ./Euler_0002 4000000
//

package main

import (
	"fmt"
	"flag"
	"strconv"
)

func main() {

	// Input
	flag.Parse()
	n, _ := strconv.Atoi(flag.Arg(0))

	// variables
	var a, b, sum = 0, 1, 0

	// Process
	for sum < n {
		b, a = a+b, b
		if b%2 == 0 {
			sum += b
		}
	}

	// Output
	fmt.Println(sum)
}
