package common

import (
	"fmt"
	"os"
	"strings"
)

func GetInput(day int) string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(string(content))
}

func GetLines(input string) []string {
	return strings.Split(input, "\n")
}
