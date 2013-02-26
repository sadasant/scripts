package main

// Pomodoro!
// References:
// -   [Wikipedia: Pomodoro Technique](https://en.wikipedia.org/wiki/Pomodoro_Technique)
// -   [List of colors for prompt](https://wiki.archlinux.org/index.php/Color_Bash_Prompt#List_of_colors_for_prompt_and_Bash)
// -   [CMD in Python and Go](http://www.darkcoding.net/software/pretty-command-line-console-output-on-unix-in-python-and-go-lang/)
// -   [time.Tick](http://golang.org/pkg/time/#example_Tick)
// -   [strconv.Itoa](http://golang.org/pkg/strconv/#Itoa)
// -   [Capture spaced user input](http://stackoverflow.com/questions/7452641/capture-spaced-user-input)
//
// It'd be cool to make it with ncurses:
// -   [ncurses with go](https://code.google.com/p/goncurses/source/browse/ncurses.go)
//

import (
	"bufio"
	"os"
	"strconv"
	"time"
)

func wait(start string, minutes int, end string) {
	var smins, ssecs string
	tchan := time.Tick(1 * time.Second)
	seconds := 0
	for _ = range tchan {
		smins = strconv.Itoa(minutes)
		if minutes < 10 {
			smins = "0" + smins
		}
		ssecs = strconv.Itoa(seconds)
		if seconds < 10 {
			ssecs = "0" + ssecs
		}
		print("\r"+start, smins+":"+ssecs, end)
		if seconds == 0 {
			if minutes == 0 {
				break
			}
			minutes--
			seconds = 60
		}
		seconds--
	}
}

const delim = '\n'

func main() {
	var task_name string
	var which_break string
	var break_time int
	bblack := "\033[1;30m"
	bwhite := "\033[1;37m"
	white := "\033[0;37m"
	bred := "\033[1;31m"
	red := "\033[0;31m"
	bgreen := "\033[1;32m"
	green := "\033[0;32m"
	end := "\033[0m"
	task_count := 1
	r := bufio.NewReader(os.Stdin)
	for {
		println("\n"+bwhite+"Pomodoro! "+white+"#"+strconv.Itoa(task_count), end)
		print("Name your task: ")
		task_name, _ = r.ReadString(delim)
		task_name = task_name[:len(task_name)-1]
		wait(red+"Waiting ", 25, " minutes for you to complete \""+bred+task_name+red+"\" :|"+end)
		print("\x07") // ASCII code 7 (BEL), or "\a"
		print("\r" + green + "You just finished \"" + bgreen + task_name + green + "\" :) \033[K\n" + end)
	ask_break:
		print("Short (s) or long (l) break?")
		which_break, _ = r.ReadString(delim)
		which_break = which_break[:len(which_break)-1]
		switch which_break {
		case "s":
			switch task_count % 4 {
			case 0:
				break_time = 15
			default:
				break_time = 3
			}
		case "l":
			switch task_count % 4 {
			case 0:
				break_time = 30
			default:
				break_time = 5
			}
		default:
			goto ask_break
		}
		wait(bblack+"Take a break of ", break_time, " minutes..."+end)
		task_count++
	}
}
