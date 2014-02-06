package main

import (
	. "fmt"
	. "math"
	"strconv"
)

func reverse(s string) string {
	l := len(s)
	r := make([]rune, l)
	for _, v := range s {
		l--
		r[l] = v
	}
	return string(r[l:])
}

func main() {
	var n int
	var p string

	Print("The largest palindrome made from the product of two numbers of length ")
	Scan(&n)

	max := int(Pow10(n) - 1)
	min := 9 * max / 10

	for i := max; ; i-- {
		for j := max; j > min; j-- {
			p = strconv.Itoa(i * j)
			if p == reverse(p) {
				Printf("is %v x %v = %s", i, j, p)
				return
			}
		}
	}

}
