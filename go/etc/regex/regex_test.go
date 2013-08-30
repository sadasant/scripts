package main

import (
	"regexp"
	"testing"
)

var test_string = "0000/00/00 00:00:00 [   ][XXXXXX] XxxxxXxxx xxxxxxxxxx xxxxxxxx"
var reg = regexp.MustCompile(`\[(.{3})]`)

func BenchmarkForLoop(b *testing.B) {
	var name string
	for i := 0; i < b.N; i++ {
		var in_brace bool
		for i := 0; i < len(test_string); i++ {
			if test_string[i] == ']' {
				break
			}
			if in_brace {
				name += string(test_string[i])
			}
			if test_string[i] == '[' {
				in_brace = true
				name = ""
			}
		}
	}
	// println(name, len(name))
}

func BenchmarkRegex(b *testing.B) {
	// var name string
	for i := 0; i < b.N; i++ {
		// name = reg.FindString(test_string)
		reg.FindString(test_string)
	}
	// println(name, len(name))
}
