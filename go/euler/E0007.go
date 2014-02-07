package main

import (
	. "fmt"
	. "math"
)

func PrimeAt(n int) int {
	if n < 3 {
		return n+1
	}
	nums := make([]bool, 3+n*int(Sqrt(float64(n))))
	c := 1
	for i, l := 3, len(nums); i < l; i+=2 {
		if !nums[i] {
			c++
			if c == n {
				return i
			}
			for j := i; j < l; j+=i {
				nums[j] = true
			}
		}
	}
	return 0
}

func main() {
	var n int

	Print("What is the prime number in position ")
	Scan(&n)

	Println("is", PrimeAt(n))
}
