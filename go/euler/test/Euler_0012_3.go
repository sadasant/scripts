// Daniel R. (sadasant.com)
// 12/02/2012
//
// Problem:
//   http://projecteuler.net/problem=10
//
// How to run:
//   ./Euler_0010 2000000
//

package main

import (
	"fmt"
	"flag"
	"strconv"
	M "math"
)

func main() {

	// Input
	flag.Parse()
	n, _ := strconv.Atoi64(flag.Arg(0))

	// Variables
	T := make([]bool, n)
	t := int64(0)
	sum := 0)
	T[1]++

	// Process
	for sum < n {
		sum = 0
		for i := 2; i < n; i++ {
			if T[i]
		}
		fmt.Println(i, T[i], I[i])
		if (T[i] == true && I[i] > n) {
			break
		}
		t += i
		T[t] = true
		i++
	}

	// Output
	fmt.Println(i, T[i], I[i])
}
