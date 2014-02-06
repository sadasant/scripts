package main

import (
	. "fmt"
	. "math"
)

func main() {
	var n, m, l int64

	Print("The largest prime factor of the number ")
	Scan(&n)

	m = 3

	for int64(Sqrt(float64(m))) < n {
		if n%m == 0 {
			n = n / m
			l = m
			m += 2
		} else {
			m += 2
		}
	}

	Println("is", l)
}
