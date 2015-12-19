package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		println("Add a regexp parameter between double quotes")
		os.Exit(2)
	}

	args := strings.Join(os.Args[1:], " ")
	println("args:", args)
	reg, err := regexp.Compile(args)
	if err != nil {
		log.Fatal(err)
	}

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "31415"
	}

	ch := make(chan string)
	go func() {
		// sudo tcpdump -vv -s 0 -XX port 31415
		err := RunCommandCh(ch, "\n", "sudo", "tcpdump", "-vv", "-s", "0", "-XX", "port", port)
		if err != nil {
			log.Fatal(err)
		}
	}()

	first_line_reg, err := regexp.Compile("\\d{2}:\\d{2}:\\d{2}.\\d{6} IP \\(tos ")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string
	var reg_matched bool
	for v := range ch {
		if first_line_reg.MatchString(v) {
			if reg_matched {
				str_lines := strings.Join(lines, "\n")
				matches := reg.FindAllStringSubmatchIndex(str_lines, -1)
				move_right := 0
				for _, k := range matches {
					k[0] += move_right
					k[1] += move_right
					first := str_lines[0:k[0]]
					last := str_lines[k[1]:]
					match := str_lines[k[0]:k[1]]
					green_match := Green(match)
					str_lines = first + green_match + last
					move_right += len(green_match) - len(match)
				}
				fmt.Println(str_lines)
				reg_matched = false
			}
			lines = []string{}
		}
		lines = append(lines, v)
		if reg.MatchString(v) {
			reg_matched = true
		}
	}
}

// RunCommandCh runs an arbitrary command and streams output to a channnel.
func RunCommandCh(stdoutCh chan<- string, cutset string, command string, flags ...string) error {
	cmd := exec.Command(command, flags...)

	output, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("RunCommand: cmd.StdoutPipe(): %v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("RunCommand: cmd.Start(): %v", err)
	}

	scanner := bufio.NewScanner(output)
	go func() {
		defer close(stdoutCh)
		for scanner.Scan() {
			text := scanner.Text()
			for {
				// Take the index of any of the given cutset
				n := strings.IndexAny(text, cutset)
				if n == -1 {
					// If not found, but still have data, send it
					if len(text) > 0 {
						stdoutCh <- text
					}
					break
				}
				// Send data up to the found cutset
				stdoutCh <- text[:n]
				// If cutset is last element, stop there.
				if n == len(text) {
					break
				}
				// Shift the text and start again.
				text = text[n+1:]
			}
		}
	}()

	if err := scanner.Err(); err != nil {
		log.Printf("RunCommand: scan process output error: %s", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("RunCommand: cmd.Wait(): %v", err)
	}

	return nil
}
