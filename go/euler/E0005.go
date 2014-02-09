package main

import . "fmt"

// Euclidean Algorithm
func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	n, lcm := 0, 1

	Print("The smallest positive number evenly divisible by all numbers from 1 to ")
	Scan(&n)

	for i := n; i > 0; i-- {
		lcm = (lcm * i) / gcd(lcm, i)
	}

	Println("is", lcm)
}
