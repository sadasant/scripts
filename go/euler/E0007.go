package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"math"
)

func AppendingPrimes(v ...int) interface{} {
	n := v[0]
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

func Sieve(v ...int) interface{} {
	n := v[0]
	if n < 3 {
		return n + 1
	}
	nums := make([]bool, 3+n*int(math.Sqrt(float64(n))))
	c := 1
	for i, l := 3, len(nums); i < l; i += 2 {
		if !nums[i] {
			c++
			if c == n {
				return i
			}
			for j := i; j < l; j += i {
				nums[j] = true
			}
		}
	}
	return 0
}

func PackagedSieve(v ...int) interface{} {
	return euler.PrimesUpTo(3+v[0]*int(math.Sqrt(float64(v[0]))))[v[0]-1]
}

func Overview(v ...int) interface{} {
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
		r := int(math.Floor(math.Sqrt(float64(n))))
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
	for count := 1; count < v[0]; c += 2 {
		if isPrime(c) {
			count++
		}
	}
	return c-2
}
 

func main() {
	euler.Init(7, "What is the Nst prime number?")
	euler.PrintTime("N = 6     | Appending primes.       Result: %v,\t\tNanoseconds: %d\n", AppendingPrimes, 6)
	euler.PrintTime("N = 6     | Sieve of Eratosthenes.  Result: %v,\t\tNanoseconds: %d\n", Sieve, 6)
	euler.PrintTime("N = 6     | With Packaged Sieve.    Result: %v,\t\tNanoseconds: %d\n", PackagedSieve, 6)
	euler.PrintTime("N = 6     | With the overview code. Result: %v,\t\tNanoseconds: %d\n", Overview, 6)
	euler.PrintTime("N = 10001 | Appending primes.       Result: %v,\tNanoseconds: %d\n", AppendingPrimes, 10001)
	euler.PrintTime("N = 10001 | Sieve of Eratosthenes.  Result: %v,\tNanoseconds: %d\n", Sieve, 10001)
	euler.PrintTime("N = 10001 | With Packaged Sieve.    Result: %v,\tNanoseconds: %d\n", PackagedSieve, 10001)
	euler.PrintTime("N = 10001 | With the overview code. Result: %v,\tNanoseconds: %d\n", Overview, 10001)
}
