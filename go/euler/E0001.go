// Daniel R. (sadasant.com)
// 21/09/2012
//
// Problem:
//   http://projecteuler.net/problem=1
//
// How to run:
//   go run E0001.go 1000
//

package main

import (
	"flag"
	"strconv"
)

func main() {

	flag.Parse()
	n, _ := strconv.Atoi(flag.Arg(0))

	var sum int
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	println(sum)
}
