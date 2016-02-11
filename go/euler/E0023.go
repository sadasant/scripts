package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"sort"
	"strings"
)

func solution(file string) (total int) {
	i := 0
	total := 0
	lower_limit := 24
	upper_limit := 28123
	// Not exactly this way
	// We need to CHECK if a number is abundant,
	// then sum it with another abundant number,
	// I guess this can go in the following ways:
	// We can... make a list of abundant numbers
	// 		for i = 0 to to i <= upper_limit/2
	// 			divisors := proper_divisors(i)
	// 			if sum(divisors) > i
	// 				abundants = append(abundants, i)
	// 				// If we sum the abundants here
	// 				// by looping over existing abundants + the new one
	// 				// we actually save some time
	// 		for i in abundants
	// 			for i in abundants
	// 				aubundants_sum = append(abundants_sum)
	//      total = sum(1 to 11)
	// 		for i = 12 to to i <= upper_limit/2
	// 			if i is not in abundants_sum
	// 				total += i
	for i < upper_limit {
		if i < lower_limit {
			total += i
		}
	}
	return
}

func main() {
	euler.Init(22, "What is the total of all the name scores in the file?")
	euler.PrintTime("For COLIN   | Result: %v, Nanoseconds: %d\n", solution, "COLIN")
	euler.PrintTime("For the file | Result: %v, Nanoseconds: %d\n", solution, euler.Download("https://projecteuler.net/project/names.txt"))
}
