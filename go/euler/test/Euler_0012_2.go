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
	pw := int64(M.Pow(float64(n), 4.0))

	// Variables
	I := make([]int64, pw)
	i := int64(1)
	T := make([]bool, pw)
	t := int64(0)

	// Process
	for i < pw {
		I[i]++
		if i < pw {
			for j := i+i; j < pw; j+=i {
				I[j]++
			}
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
