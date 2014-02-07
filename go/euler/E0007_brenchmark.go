package main

import (
	"time"
	. "math"
)

func E0007(n int) int {
	primes := []int{3,5,7,11,13}
	c := len(primes)
	i := primes[c-1]+2
L:
	for ; c < n; i+=2 {
		for _, v := range primes {
			if i%v == 0 {
				continue L
			}
		}
		primes = append(primes, i)
		c++
	}
	return primes[n-2]
}

func Boolean(n int) int {
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

func Overview(n int) int {
	isPrime := func(n int) bool {
		switch {
		case n==1:
			return false
		case n<4:
			return true
		case n%2 == 0:
			return false
		case n<9:
			return true
		case n%3 == 0:
			return false
		}
		r := int(Floor(Sqrt(float64(n))))
		f := 5
		for f <= r {
			if n%f == 0 || n%(f+2) == 0 {
				return false
			}
			f += 6
		}
		return true
	}
	c := 3
	for count := 1; count < n; c += 2 {
		if isPrime(c) {
			count++
		}
	}
	return c-2
}

func benchmark(n int) {
	ta := time.Now()
	a := E0007(n)
	tb := time.Now()
	b := Overview(n)
	tc := time.Now()
	c := Boolean(n)
	t3 := time.Now()
	println("E0007   ", tb.Sub(ta), a)
	println("Overview", tc.Sub(tb), b)
	println("Boolean ", t3.Sub(tc), c)
}

func main() {
	benchmark(6)
	benchmark(10001)
}
