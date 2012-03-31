// Daniel R. (sadasant.com)
// 17/02/2012
//
// Problem:
//   http://projecteuler.net/problem=12
//
// How to run:
//   ./Euler_0012 500
//
// Why don't we use Euler_10 ?
//

package main

import (
	"fmt"
	"flag"
	"strconv"
)

// Funcs
func countDivisors(n int) int {
	if n < 2 { return n }
	max := n/2
	c := 1
	for i := 1; i <= max; i++ {
		if n%i == 0 { c++ }
	}
	return c
}

func main() {

	// Input
	flag.Parse()
	n, _ := strconv.Atoi(flag.Arg(0))

	// Variables
	max := 0
	i := 1
	t := 0

	// Process
	for max < n {
		t += i
		i++
		divs := countDivisors(t)
		if divs > max {
			max = divs
			fmt.Println(t, max)
		}
	}

	// Output
	fmt.Println(t, max)
}
