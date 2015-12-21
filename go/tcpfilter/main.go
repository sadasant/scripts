package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		println("Add a regexp query")
		os.Exit(2)
	}

	args := strings.Join(os.Args[1:], " ")
	println("Filtering by:", args)
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
	hex_byte_reg, err := regexp.Compile("[0-9abcdef]{4}(\\s|$)")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string
	var reg_matched bool
	for v := range ch {
		if first_line_reg.MatchString(v) {
			str_lines := strings.Join(lines, "\n")
			hex_bytes := hex_byte_reg.FindAllString(str_lines, -1)
			hex_content := strings.Join(strings.Split(strings.Join(hex_bytes, ""), " "), "")
			content, err := hex.DecodeString(hex_content)
			str_content := onlyNiceCaracters(string(content))
			if err != nil {
				log.Fatal(err)
			}
			if reg_matched {
				indexes := reg.FindAllStringSubmatchIndex(str_lines, -1)
				str_lines = Paint(str_lines, indexes, Green)
				fmt.Println(str_lines)
			}
			if reg.MatchString(str_content) {
				if !reg_matched {
					fmt.Println(str_lines)
				}
				println("Hex content:", hex_content)
				indexes := reg.FindAllStringSubmatchIndex(str_content, -1)
				str_content = Paint(str_content, indexes, Green)
				println(fmt.Sprintf("Content MATCHED: %s", str_content))
				println()
			} else if reg_matched {
				println("Hex content:", hex_content)
				println(fmt.Sprintf("Content: %s", str_content))
				println()
			}
			// Reset
			reg_matched = false
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
