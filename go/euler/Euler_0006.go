// Daniel R. (sadasant.com)
// 11/02/2012
//
// Problem:
//   http://projecteuler.net/problem=6
//
// How to run:
//   ./Euler_0006 100
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
	n, _ := strconv.Atof64(flag.Arg(0))

	// Variables
	sum_sq, sq_sum := 0.0, 0.0

	// Process
	for i := 1.0; i <= n; i += 1 {
		sum_sq += math.Pow(i, 2.0)
		sq_sum += i
	}
	sq_sum = math.Pow(sq_sum, 2.0)

	// Output
	fmt.Println(int(sq_sum - sum_sq))
}
