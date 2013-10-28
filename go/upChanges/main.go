package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	netrc "code.google.com/p/go-netrc/netrc"
)

var version string = "v0.0.0"

func main() {
	machine := *flag.String("machine", "", "Remote machine name")
	login := *flag.String("login", "", "User of the remote machine")
	password := *flag.String("password", "", "Password of the remote machine")
	directory := *flag.String("directory", "", "Password of the remote machine")

	flag.Parse()

	println("")
	println("upChanges", version)
	println("")

	if machine == "" {
		if len(os.Args) > 1 {
			machine = os.Args[1]
		} else {
			exit("No machine given.")
		}
	}

	if directory == "" {
		if len(os.Args) > 2 {
			directory = os.Args[len(os.Args)-1]
		} else {
			exit("No directory given.")
		}
	}

	fmt.Println(machine)
	fmt.Println(directory)

	usr, err := user.Current()
	exitIf(err)

	if login == "" && password == "" {
		// Open .netrc
		mach, err := netrc.FindMachine(usr.HomeDir+"/.netrc", machine)
		exitIf(err)
		fmt.Printf("%v", mach)
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func exitIf(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
