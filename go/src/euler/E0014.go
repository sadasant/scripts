// Euler 14 http://projecteuler.net/problem=14
// by Daniel Rodríguez http://sadasant.com

package main

import "fmt"

func main() {
  fmt.Println("Hello, playground")
	var longest, chain int
	var longest_n int64
	var n int64 = 1
	for ; n < 1000000; n++ {
		chain = Collatz(n)
		if chain > longest {
			longest = chain
			longest_n = n
		}
	        // fmt.Println(n, chain, longest)
	}
	fmt.Println(longest_n)
	return
}

func Collatz(n int64) int {
	var chain int
L:
	chain++
	switch {
	case n == 1:
		return chain
	case n%2 == 0:
		n = n / 2
	default:
		n = 3*n + 1
	}
	goto L
}
