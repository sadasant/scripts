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
		println("Use as: tcpfilter -tcpdump [tcpdump options] -filter [regexp]")
		os.Exit(2)
	}

	args := strings.Join(os.Args[1:], " ")
	println("args:", args)

	tcpdump_params := "tcpdump -vv -s 0 -XX"
	if matched, err := regexp.MatchString("-tcpdump", args); matched && err == nil {
		tcpdump_params = "tcpdump" + strings.Split(args, "-tcpdump")[1]
		// Removing -filter from tcpdump_params
		if matched, err := regexp.MatchString("-filter", tcpdump_params); matched && err == nil {
			tcpdump_params = strings.Split(tcpdump_params, "-filter")[0]
		}
	}
	println(Bold("Running:"), tcpdump_params)

	filter_param := ".*"
	if matched, err := regexp.MatchString("-filter", args); matched && err == nil {
		filter_param = strings.Split(args, "-filter")[1][1:]
	}
	println(Bold("Filtering by:"), filter_param)

	filter, err := regexp.Compile(filter_param)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan string)
	go func() {
		err := RunCommandCh(ch, "\n", "sudo", strings.Split(tcpdump_params, " ")...)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Match the first line of the tcpdump result
	first_line_regexp, err := regexp.Compile("\\d{2}:\\d{2}:\\d{2}.\\d{6} IP \\(tos ")
	if err != nil {
		log.Fatal(err)
	}

	// Match all the hex bytes of the content of the packages
	hex_byte_regexp, err := regexp.Compile("[0-9abcdef]{4}(\\s|$)")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	var filter_matched bool
	for v := range ch {

		if first_line_regexp.MatchString(v) {
			str_lines := strings.Join(lines, "\n")

			hex_bytes := hex_byte_regexp.FindAllString(str_lines, -1)
			hex_content := strings.Join(strings.Split(strings.Join(hex_bytes, ""), " "), "")
			content, hex_err := hex.DecodeString(hex_content)
			str_content := onlyNiceCaracters(string(content))

			if filter_matched {
				indexes := filter.FindAllStringSubmatchIndex(str_lines, -1)
				str_lines = Paint(str_lines, indexes, Green)
				fmt.Println(str_lines)
			}

			if hex_err == nil {
				if hex_err == nil && filter.MatchString(str_content) {
					if !filter_matched {
						fmt.Println(str_lines)
					}
					println("Hex content:", hex_content)
					indexes := filter.FindAllStringSubmatchIndex(str_content, -1)
					str_content = Paint(str_content, indexes, Green)
					println(fmt.Sprintf("Content MATCHED: %s", str_content))
					println()
				} else if filter_matched {
					println("Hex content:", hex_content)
					println(fmt.Sprintf("Content: %s", str_content))
					println()
				}
			} else if filter_matched {
				println()
			}

			// Reset
			filter_matched = false
			lines = []string{}
		}

		lines = append(lines, v)
		if filter.MatchString(v) {
			filter_matched = true
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
