// Daniel R. (sadasant.com)
// 01/02/2012
//
// Problem:
//   http://projecteuler.net/problem=1
//
// How to run:
//   ./Euler_0001 1000
//

package main

import (
	"fmt"
	"flag"
	"strconv"
)

func main() {

	// input
	flag.Parse()
	n, _ := strconv.Atoi(flag.Arg(0))

	// vars
	var sum int

	// process
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	// output
	fmt.Println(sum)
}

