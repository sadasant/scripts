// This is an example of the talk:
//
// 	10 things you (probably) don't know about Go
//
// Specifically from the page 10th slide:
//
// 	http://talks.golang.org/2012/10things.slide#10
//
// The comments inside the code are extracts from the paper:
//
// 	Concurrent Programming in Go
//
// Which can be found here: http://www.cs.ucla.edu/~rohitkk/go_manual.pdf
//
// Thanks.
//

package main

import "fmt"

var battle = make(chan string)

func warrior(name string, done chan bool) {
	// The select statement is used to choose
	// a send or receive from a set of channels.
	// The one which can proceed is selected and excecuted.
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
