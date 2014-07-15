package main

import "github.com/sadasant/scripts/go/euler/euler"

// From: http://www.mathblog.dk/project-euler-21-sum-of-amicable-pairs/
func sumOfPrimeFactors(number int, primes *[]int) int {
	var sum float32 = 1
	var j float32
	var i int
	n := float32(number)
	p := float32((*primes)[0])
	l := len(*primes)
	for p * p <= n && n > 1 && i < l {
		p = float32((*primes)[i])
		i++
		if int(n) % int(p) == 0 {
			j = p * p
			n /= p
			for int(n) % int(p) == 0 {
				j *= p
				n /= p
			}
			sum *= (j-1)/(p-1)
		}
	}
	if n > 1 {
		sum *= n+1
	}
	return int(sum) - number
}

func solution(n int) int {
	var sum int
	primes := euler.PrimesUpTo(n)
	dict := map[int]int{}
	for i := 2; i <= n; i++ {
		factors := sumOfPrimeFactors(i, &primes)
		if factors > i {
			dict[i] = factors
		} else if factors < i && dict[factors] == i {
			sum += i + factors;
		}
	}
	return sum
}

func main() {
	euler.Init(21, "Evaluate the sum of all the amicable numbers under...")
	euler.PrintTime("Under    220 | Result: %v, Nanoseconds: %d\n", solution, 220)
	euler.PrintTime("Under    284 | Result: %v, Nanoseconds: %d\n", solution, 284)
	euler.PrintTime("Under  10000 | Result: %v, Nanoseconds: %d\n", solution, 10000)
	euler.PrintTime("Under 100000 | Result: %v, Nanoseconds: %d\n", solution, 100000)
}
