package main

const color_bold = "\x1b[1m"
const color_green = "\x1b[32m"
const color_normal = "\x1b[0m"

func Green(s string) string {
	return color_green + s + color_normal
}

func Bold(s string) string {
	return color_bold + s + color_normal
}

func Paint(str string, indexes [][]int, color func(s string) string) string {
	move_right := 0
	for _, k := range indexes {
		k[0] += move_right
		k[1] += move_right
		first := str[0:k[0]]
		last := str[k[1]:]
		match := str[k[0]:k[1]]
		color_match := color(match)
		str = first + color_match + last
		move_right += len(color_match) - len(match)
	}
	return str
}

func onlyNiceCaracters(s string) string {
	var new_s string
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 126 {
			new_s += "."
		} else {
			new_s += string(s[i])
		}
	}
	return new_s
}
