package main


import (
	"github.com/sadasant/scripts/go/euler/euler"
	"math"
)

func solution(v ...int) interface{} {
	n := float64(v[0])
	var i, r float64 = 1, 1
	for ; i <= n; i++ {
		r *= (n+i)/i
	}
	return int(math.Ceil(r))
}

func main() {
	euler.Init(15, "How many such routes are there through a 20×20 grid?")
	euler.PrintTime("Combinatorial Solution. Result: %v, Nanoseconds: %d\n", solution, 20)
}
