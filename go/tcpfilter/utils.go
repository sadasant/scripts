package main

const color_green = "\x1b[32m"
const color_normal = "\x1b[0m"

func Green(s string) string {
	return color_green + s + color_normal
}
