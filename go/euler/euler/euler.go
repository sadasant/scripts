package euler

import (
	"fmt"
	"time"
)

const (
	b_black   = "\x1b[30;1m"
	b_red     = "\x1b[31;1m"
	b_green   = "\x1b[32;1m"
	b_yellow  = "\x1b[33;1m"
	b_blue    = "\x1b[34;1m"
	b_magenta = "\x1b[35;1m"
	b_cian    = "\x1b[36;1m"
	b_white   = "\x1b[37;1m"
	black     = "\x1b[30m"
	red       = "\x1b[31m"
	green     = "\x1b[32m"
	yellow    = "\x1b[33m"
	blue      = "\x1b[34m"
	magenta   = "\x1b[35m"
	cian      = "\x1b[36m"
	white     = "\x1b[37m"
	normal    = "\x1b[0m"
)

func Init(n int, text string) {
	fmt.Printf("%vProject Euler, problem %v%v%v\n%s%v\n", b_yellow, n, normal, yellow, text, normal)
}

func Time(f func(v ...int) int, v ...int) (int, time.Duration) {
	ta := time.Now()
	r := f(v...)
	tb := time.Now()
	return r, tb.Sub(ta)
}

func PrintTime(msg string, f func(v ...int) int, v ...int) {
	r, t := Time(f, v...)
	fmt.Printf(msg, r, t)
}

func Sequence(v ...int) []int {
	var start, steps, end int
	start = v[0]
	switch {
	case len(v) < 3:
		end = v[1]
		steps = 1
		break
	default:
		end = v[2]
		steps = v[1]
		break
	}
	r := make([]int, end+1-start)
	for i := start; i <= end; i += steps {
		r[i-start] = i
	}
	return r
}
