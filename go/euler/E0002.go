package main

import . "fmt"

func main() {

	n, a, b, s := 0, 0, 1, 0

	Println("Sum the Fibonacci sequence's even terms lower than:")
	Scan(&n)

	for b < n {
		b, a = a+b, b
		if b%2 == 0 {
			s += b
		}
	}

	Println(s)
}
