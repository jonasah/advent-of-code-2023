package day13

import (
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
)

type Pattern struct {
	data          []string
	width, height int
}

func (p *Pattern) row(r int) string {
	return p.data[r]
}

func (p *Pattern) col(c int) string {
	var sb strings.Builder
	for _, row := range p.data {
		sb.WriteByte(row[c])
	}

	return sb.String()
}

func Part1(input string) int {
	patterns := parseInput(input)

	numCols := 0
	numRows := 0
	for _, pattern := range patterns {
		row := findReflection(pattern.height, pattern.row, false, -1)

		if row != -1 {
			numRows += row + 1
		} else {
			numCols += findReflection(pattern.width, pattern.col, false, -1) + 1
		}
	}

	return numCols + 100*numRows
}

func Part2(input string) int {
	patterns := parseInput(input)

	numCols := 0
	numRows := 0
	for _, pattern := range patterns {
		prevRow := findReflection(pattern.height, pattern.row, false, -1)
		row := findReflection(pattern.height, pattern.row, true, prevRow)

		if row != -1 && row != prevRow {
			numRows += row + 1
			continue
		}

		prevCol := findReflection(pattern.width, pattern.col, false, -1)
		col := findReflection(pattern.width, pattern.col, true, prevCol)

		if col != -1 && col != prevCol {
			numCols += col + 1
		}
	}

	return numCols + 100*numRows
}

func parseInput(input string) []Pattern {
	patterns := make([]Pattern, 0)
	for _, p := range strings.Split(input, "\n\n") {
		lines := common.GetLines(p)
		pattern := Pattern{width: len(lines[0]), height: len(lines)}
		for _, l := range lines {
			pattern.data = append(pattern.data, l)
		}

		patterns = append(patterns, pattern)
	}

	return patterns
}

func findReflection(size int, getSequence func(i int) string, smudge bool, skip int) int {
	maxDiff := 0
	if smudge {
		maxDiff = 1
	}

outer:
	for i := 0; i < size-1; i++ {
		diff := 0
		for lower, upper := i, i+1; lower >= 0 && upper < size; lower, upper = lower-1, upper+1 {
			diff += countDifference(getSequence(lower), getSequence(upper))

			if diff > maxDiff {
				continue outer
			}
		}

		if i != skip {
			return i
		}
	}

	return -1
}

func countDifference(s1, s2 string) int {
	n := 0
	for i, c := range s1 {
		if c != rune(s2[i]) {
			n++
		}
	}

	return n
}
