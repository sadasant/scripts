package main

import "github.com/sadasant/scripts/go/euler/euler"

var months_limit = []int {
	31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

func solution(min_year, max_year int) interface{} {
	var week_day  = 2;
	var month_day = 1;
	var month = 0;
	var year  = 1900;
	var saturday_firsts = 0;
	for year <= max_year{
		if year >= min_year && week_day == 1 && month_day == 1 {
			// euler.Println(year, month, month_day, week_day)
			saturday_firsts++;
		}

		if week_day < 7 {
			week_day++
		} else {
			week_day = 1
		}

		if month == 1 && month_day == 29 && (year%100 == 0 && year%400 == 0 || year%4 == 0) {
			month_day = 1
			month++
			year++
		} else if month_day >= 28 && month_day == months_limit[month] {
			month_day = 1
			if month < 11 {
				month++
			} else {
				month = 0
				year++
			}
		} else {
			month_day++
		}
	}
	return saturday_firsts
}

func solution2(years ...int) int {
	day, count := 2, 0
	for _, y := range years {
		for m := 0; m < 12; m++ {
			if m == 1 && y%4 == 0 {
				day += 29
			} else {
				day += months_limit[m]
			}
			if day%7 == 0 {
				count++
			}
		}
	}
	return count
}

func solution3(n_years int) int {
	return 12*n_years/7
}

func main() {
	euler.Init(19, "How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?")
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution, 1901, 2000)
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution2, euler.Sequence(1901, 2000)...)
	euler.PrintTime("Result: %v, Nanoseconds: %d\n", solution3, 100)
}
