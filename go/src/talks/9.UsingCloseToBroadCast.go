package main

import (
	"fmt"
	"math/rand"
	"time"
)

func waiter(i int, block, done chan bool) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Println(i, "waiting...")
	<-block
	fmt.Println(i, "done!")
	done <- true
}

func main() {
	block, done := make(chan bool), make(chan bool)
	for i := 0; i < 4; i++ {
		go waiter(i, block, done)
	}
	time.Sleep(5 * time.Second)
	close(block)
	for i := 0; i < 4; i++ {
		<-done
	}
}
