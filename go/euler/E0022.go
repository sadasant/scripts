package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"sort"
	"strings"
)

func solution(file string) (total int) {
	lines := strings.Split(file, ",")
	sort.Strings(lines)
	lower_limit := 64
	for k, v := range lines {
		worth := 0
		// euler.Printf("%v %v ", k, v)
		for _, _v := range v {
			i_v := int(_v)
			if i_v > lower_limit {
				// euler.Printf("%s=%v ", string(_v), i_v - lower_limit)
				worth += i_v - lower_limit
			}
		}
		total += (k+1) * worth
		// euler.Printf("\n")
	}
	return
}

func main() {
	euler.Init(22, "What is the total of all the name scores in the file?")
	euler.PrintTime("For COLIN   | Result: %v, Nanoseconds: %d\n", solution, "COLIN")
	euler.PrintTime("For the file | Result: %v, Nanoseconds: %d\n", solution, euler.Download("https://projecteuler.net/project/names.txt"))
}
