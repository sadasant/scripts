package main

import (
	"github.com/sadasant/scripts/go/euler/euler"
	"strconv"
)

var under_ten [9]string = [9]string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

var under_twenty [9]string = [9]string{
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

var tens [9]string = [9]string{
	"ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

func solution(N ...int) interface{} {
	s := 0
	for _, n := range N {
		ns := strconv.Itoa(n)
		euler.Print(ns, " ")
		l := len(ns)
		w := ""
		for i := 0; i < l; i++ {
			in := int(ns[i]-'0')
			if in == 0 {
				continue
			}
			switch l-1-i {
			case 3:
				w = under_ten[in-1] + "thousand"
				break
			case 2:
				w = under_ten[in-1] + "hundred"
				break
			case 1:
				inext := int(ns[i+1]-'0')
				if ns[i] == '1' && inext != 0 {
					w = under_twenty[inext-1]
				} else {
					w = tens[in-1]
				}
				if ns[i] == '1' {
					i++
				}
				if l > 2 {
					w = "and"+w
				}
				break
			default:
				w = under_ten[in-1]
				if l > 2 && int(ns[i-1]-'0') == 0 {
					w = "and"+w
				}
			}
			euler.Print(w)
			s += len(w)
		}
		euler.Print("\n")
	}
	return s
}

func main() {
	euler.Init(17, "If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?")
	euler.PrintTime("N = 342 | Result: %v, Nanoseconds: %d\n", solution, 342)
	euler.PrintTime("N = 115 | Result: %v, Nanoseconds: %d\n", solution, 115)
	euler.PrintTime("N = 1000 | Result: %v, Nanoseconds: %d\n", solution, euler.Sequence(1, 1000)...)
}
