package main

import . "fmt"

func main() {

	var n, i, s int

	Print("The sum all the multiples of 3 or 5 below ")
	Scan(&n)

	for ; i < n; i++ {
		if i%3*i%5 == 0 {
			s += i
		}
	}

	Println("is", s)
}
