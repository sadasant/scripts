// Daniel R. (sadasant.com)
// 22/09/2012
//
// Problem:
//   http://projecteuler.net/problem=4
//
// How to run:
//   go run E0004.go 3
//

package main

import (
	"flag"
	"math"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i, j = i+1, j-1
	}
	return string(runes)
}

func main() {
	flag.Parse()
	n, _ := strconv.Atoi(flag.Arg(0))

	max := int(math.Pow10(n))
	min := max - int(math.Pow10(n-1))
	r0, r1, largest_prod := 0, 0, 0

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

	println(r0, r1, largest_prod)
}
