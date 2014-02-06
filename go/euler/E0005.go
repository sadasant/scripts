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
	var n int

	Print("The smallest positive number evenly divisible by all numbers from 1 to ")
	Scan(&n)

	mults := make([]int, n)
	lcm := 1

	for i := n; i > 0; i-- {
		for j, v := range mults {
			if v == 0 {
				mults[j] = i
				lcm = (lcm * i) / gcd(lcm, i)
				break
			} else if v%i == 0 {
				Println("is", lcm)
				return
			}
		}
	}

}
