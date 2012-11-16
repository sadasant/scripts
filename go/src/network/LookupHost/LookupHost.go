package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "hostname")
		return
	}

	name := os.Args[1]
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Resolution Error", err)
		return
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
}
