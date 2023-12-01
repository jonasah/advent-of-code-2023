package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/jonasah/advent-of-code-2023/lib/common"
)

func Part1(input string) int {
	lines := common.GetLines(input)

	codes := make([]int, 0, len(lines))

	for _, line := range lines {
		first := strings.IndexFunc(line, unicode.IsDigit)
		last := strings.LastIndexFunc(line, unicode.IsDigit)

		code, _ := strconv.Atoi(fmt.Sprintf("%v%v", line[first]-'0', line[last]-'0'))
		codes = append(codes, code)
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
		fwd := line
		i := 0

	loop:
		for i < len(fwd) {
			for a, n := range numbers {
				if strings.HasPrefix(fwd[i:], n) {
					fwd = fwd[:i] + strings.Replace(fwd[i:], n, fmt.Sprintf("%d", a+1), 1)
					break loop
				}
			}

			i++
		}

		back := line
		i = len(back) - 1

	loop2:
		for i >= 0 {
			for a, n := range numbers {
				if strings.HasPrefix(back[i:], n) {
					back = back[:i] + strings.Replace(back[i:], n, fmt.Sprintf("%d", a+1), 1)
					break loop2
				}
			}

			i--
		}

		first := strings.IndexFunc(fwd, unicode.IsDigit)
		last := strings.LastIndexFunc(back, unicode.IsDigit)

		code, _ := strconv.Atoi(fmt.Sprintf("%v%v", fwd[first]-'0', back[last]-'0'))
		codes = append(codes, code)

		fmt.Println(line, fwd, back, code)
	}

	return sum(codes)
}

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
