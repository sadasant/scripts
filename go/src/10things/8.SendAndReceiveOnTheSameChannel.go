package main

import "fmt"

var battle = make(chan string)

func warrior(name string, done chan bool) {
	select {
	case opponent := <-battle:
		fmt.Printf("%s beat %s\n", name, opponent)
	case battle <- name:
		// I lost :-(
	}
	// Each of this goroutines
	// sends a value `true` to the channel `done`
	// after it has finished.
	done <- true
}

func main() {
	done := make(chan bool)
	langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}
	for _, l := range langs {
		go warrior(l, done)
	}
	// In the main goroutine
	// we wait for all child goroutines to finish
	// by receiving the values of the channels
	// and discarding them.
	for _ = range langs {
		<-done
	}
	// Here all the spawned goroutines have finished.
}
