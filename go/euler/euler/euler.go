package euler

import (
	"fmt"
	"time"
	"reflect"
	"flag"
)

var verbose = flag.Bool("v", false, "Show euler.Prints")

func init() {
	flag.Parse()
}

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

func Time(f interface{}, v ...interface{}) (interface{}, time.Duration) {
	vf := reflect.ValueOf(f)
	var vv = make([]reflect.Value, len(v))
	for i, iv := range v {
		vv[i] = reflect.ValueOf(iv)
	}
	ta := time.Now()
	r := vf.Call(vv)[0].Interface()
	tb := time.Now()
	return r, tb.Sub(ta)
}

func PrintTime(msg string, f interface{}, v ...interface{}) {
	r, t := Time(f, v...)
	fmt.Printf(msg, r, t)
}

func Sequence(v ...int) []interface{} {
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
	r := make([]interface{}, end+1-start)
	for i := start; i <= end; i += steps {
		r[i-start] = i
	}
	return r
}

func Print(v ...interface{}) {
	if *verbose {
		fmt.Print(v...)
	}
}

func Println(v ...interface{}) {
	if *verbose {
		fmt.Println(v...)
	}
}

func Printf(s string, v ...interface{}) {
	if *verbose {
		fmt.Printf(s, v...)
	}
}
