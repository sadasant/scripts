package main

import . "fmt"

func main() {
	var n int

	Print("What is the prime number ")
	Scan(&n)

	primes := make([]int, n)

	copy(primes, []int{2,3,5})

	for l, i := 3, 7; l < n; i+=2 {
		for j, v := range primes {
			if v == 0 {
				primes[j] = i
				l++
				break
			}
			if i%v == 0 {
				break
			}
		}
	}

	Println("is", primes[n-1])
}
