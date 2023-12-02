package day1

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/jonasah/advent-of-code-2023/lib/common"
)

func Part1(input string) int {
	lines := common.GetLines(input)

	codes := make([]int, 0, len(lines))

	for _, line := range lines {
		codes = append(codes, getCode(line))
	}

	return sum(codes)
}

func Part2(input string) int {
	lines := common.GetLines(input)

	codes := make([]int, 0, len(lines))

	var numbers = []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	for _, line := range lines {
		for i, n := range numbers {
			line = strings.ReplaceAll(line, n, fmt.Sprintf("%s%d%s", n, i+1, n[len(n)-1:]))
		}

		codes = append(codes, getCode(line))
	}

	return sum(codes)
}

func getCode(line string) int {
	first := strings.IndexFunc(line, unicode.IsDigit)
	last := strings.LastIndexFunc(line, unicode.IsDigit)

	return int((line[first]-'0')*10 + line[last] - '0')
}

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
