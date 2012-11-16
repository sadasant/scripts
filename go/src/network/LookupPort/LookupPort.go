package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type service\n", os.Args[0])
		return
	}

	netType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(netType, service)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println("Service port ", port)
}
