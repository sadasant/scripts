package main

import "github.com/sadasant/scripts/go/euler/euler"

// From: http://www.mathblog.dk/project-euler-21-sum-of-amicable-pairs/
func sumOfPrimeFactors(number int, primes *[]int) int {
	sum := 1
	var j, i int
	n := number
	p := (*primes)[0]
	l := len(*primes)
	for p*p <= n && n > 1 && i < l {
		p = (*primes)[i]
		i++
		if n%p == 0 {
			j = p * p
			n /= p
			for n%p == 0 {
				j *= p
				n /= p
			}
			sum *= (j - 1) / (p - 1)
		}
	}
	if n > 1 {
		sum *= n + 1
	}
	return sum - number
}

func solution(n int, primes *[]int) int {
	var sum int
	dict := map[int]int{}
	for i := 2; i <= n; i++ {
		factors := sumOfPrimeFactors(i, primes)
		if factors > i {
			dict[i] = factors
		} else if factors < i && dict[factors] == i {
			sum += i + factors
		}
	}
	return sum
}

func main() {
	primes := euler.PrimesUpTo(100000)
	euler.Init(21, "Evaluate the sum of all the amicable numbers under...")
	euler.PrintTime("Under    220 | Result: %v, Nanoseconds: %d\n", solution, 220, &primes)
	euler.PrintTime("Under    284 | Result: %v, Nanoseconds: %d\n", solution, 284, &primes)
	euler.PrintTime("Under  10000 | Result: %v, Nanoseconds: %d\n", solution, 10000, &primes)
	euler.PrintTime("Under 100000 | Result: %v, Nanoseconds: %d\n", solution, 100000, &primes)
}
