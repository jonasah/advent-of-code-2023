package day9

import (
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/sliceconv"
)

func Solve(input string, part int) int {
	lines := common.GetLines(input)

	sum := 0
	for _, line := range lines {
		values := sliceconv.Atoi(strings.Fields(line))
		sum += findNextValue(values, part)
	}

	return sum
}

func findNextValue(values []int, part int) int {
	if isAllZeroes(values) {
		return 0
	}

	if part == 1 {
		return values[len(values)-1] + findNextValue(diff(values), part)
	}

	return values[0] - findNextValue(diff(values), part)
}

func diff(values []int) []int {
	d := make([]int, 0, len(values)-1)
	for i, v := range values[1:] {
		d = append(d, v-values[i])
	}

	return d
}

func isAllZeroes(values []int) bool {
	return !slices.ContainsFunc(values, func(i int) bool { return i != 0 })
}
