package color

import "fmt"

const escape = "\x1b"

const (
	none = iota
	red
	green
	yellow
	blue
	purple
)

func color(c int) string {
	if c == none {
		return fmt.Sprintf("%s[%dm", escape, c)
	}

	return fmt.Sprintf("%s[3%dm", escape, c)
}

func format(c int, text string) string {
	return color(c) + text + color(none)
}

func Red(text string) string {
	return format(red, text)
}

func Green(text string) string {
	return format(green, text)
}

func Yellow(text string) string {
	return format(yellow, text)
}

func Blue(text string) string {
	return format(blue, text)
}

func Purple(text string) string {
	return format(purple, text)
}
