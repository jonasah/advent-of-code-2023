package day3

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
)

type Point struct {
	r, c int
}

type Number struct {
	n int
	p Point
}

type Symbol struct {
	s rune
	p Point
}

func (n Number) isAdjacent(s Symbol) bool {
	l := len(strconv.Itoa(n.n))
	return math.Abs(n.p.r-s.p.r) <= 1 && s.p.c >= n.p.c-1 && s.p.c <= n.p.c+l
}

func Part1(input string) int {
	numbers, symbols := parseInput(input, false)

	sum := 0
outer:
	for _, n := range numbers {
		for _, s := range symbols {
			if n.isAdjacent(s) {
				sum += n.n
				continue outer
			}
		}
	}

	return sum
}

func Part2(input string) int {
	numbers, gears := parseInput(input, true)

	sum := 0
	for _, g := range gears {
		adjNumbers := make([]Number, 0)
		for _, n := range numbers {
			if n.isAdjacent(g) {
				adjNumbers = append(adjNumbers, n)
			}
		}

		if len(adjNumbers) == 2 {
			sum += adjNumbers[0].n * adjNumbers[1].n
		}
	}

	return sum
}

func parseInput(input string, gearsOnly bool) ([]Number, []Symbol) {
	lines := common.GetLines(input)

	numbers := make([]Number, 0)
	symbols := make([]Symbol, 0)

	for r, line := range lines {
		c := 0

		for len(line) > 0 {
			firstChar := rune(line[0])

			if firstChar == '.' {
				line = line[1:]
				c++
				continue
			}

			if unicode.IsDigit(firstChar) {
				i := strings.IndexFunc(line, func(r rune) bool { return !unicode.IsDigit(r) })
				if i == -1 {
					i = len(line)
				}
				n, _ := strconv.Atoi(line[:i])
				numbers = append(numbers, Number{n, Point{r, c}})
				line = line[i:]
				c += i
				continue
			}

			if !gearsOnly || (gearsOnly && firstChar == '*') {
				symbols = append(symbols, Symbol{firstChar, Point{r, c}})
			}
			line = line[1:]
			c++
		}
	}

	return numbers, symbols
}
