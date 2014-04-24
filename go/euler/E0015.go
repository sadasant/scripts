package main


import (
	"github.com/sadasant/scripts/go/euler/euler"
	"math"
)

func solution(n int) interface{} {
	n64 := float64(n)
	var i, r float64 = 1, 1
	for ; i <= n64; i++ {
		r *= (n64+i)/i
	}
	return int(math.Ceil(r))
}

func main() {
	euler.Init(15, "How many such routes are there through a 20×20 grid?")
	euler.PrintTime("Combinatorial Solution. Result: %v, Nanoseconds: %d\n", solution, 20)
}
