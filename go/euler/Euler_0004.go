// Daniel R. (sadasant.com)
// 04/02/2012
//
// Problem:
//   http://projecteuler.net/problem=4
//
// How to run:
//   ./Euler_0004 3
//

package main

import (
	"fmt"
	"flag"
	"strconv"
	"math"
	"utf8"
)

func reverse(s string) string {
	o := make([]int, utf8.RuneCountInString(s))
	i := len(o)
	for _, c := range s {
		i--
		o[i] = c
	}
	return string(o)
}

func main() {

	// Input
	flag.Parse()
	n, _ := strconv.Atoi(flag.Arg(0))

	// Variables
	max := int(math.Pow10(n))
	min := max - int(math.Pow10(n-1))
	r0, r1, largest_prod := 0, 0, 0

	// Process
	for i := max; i > min; i-- {
		for j := max; j > min; j-- {
			prod := i * j
			str_prod := strconv.Itoa(i * j)
			if str_prod == reverse(str_prod) && prod > largest_prod {
				r0 = i
				r1 = j
				largest_prod = prod
			}
		}
	}

	// Output
	fmt.Println(r0, r1, largest_prod)
}
