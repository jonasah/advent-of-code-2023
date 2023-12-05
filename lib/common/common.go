package common

import (
	"fmt"
	"os"
	"strconv"
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

func ToInts(a []string) []int {
	l := make([]int, 0, len(a))
	for _, s := range a {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		l = append(l, i)
	}

	return l
}
