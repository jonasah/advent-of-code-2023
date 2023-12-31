package day21

import (
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
)

const (
	start = 'S'
	plot  = '.'
	rock  = '#'

	right = 1
	down  = 1i
	left  = -1
	up    = -1i
)

type Position complex128

func (p Position) x() int {
	return int(real(p))
}

func (p Position) y() int {
	return int(imag(p))
}

type Garden struct {
	data          [][]rune
	width, height int
}

func (g Garden) at(p Position) rune {
	return g.data[p.y()][p.x()]
}

func Part1(input string, steps int) int {
	garden, start := parseInput(input)

	dir := []Position{right, down, left, up}
	positions := []Position{start}

	for steps > 0 {
		var newPositions []Position

		for _, pos := range positions {
			for _, d := range dir {
				p := pos + d

				if garden.at(p) == rock {
					continue
				}

				if slices.Contains(newPositions, p) {
					continue
				}

				newPositions = append(newPositions, p)
			}
		}

		positions = newPositions
		steps--
	}

	return len(positions)
}

func Part2(input string) int {
	return 0
}

func parseInput(input string) (Garden, Position) {
	lines := common.GetLines(input)

	garden := Garden{width: len(lines[0]), height: len(lines)}
	var startPos Position
	for r, line := range lines {
		garden.data = append(garden.data, []rune(line))

		idx := strings.IndexRune(line, start)
		if idx != -1 {
			startPos = Position(complex(float64(idx), float64(r)))
			garden.data[r][idx] = plot
		}
	}

	return garden, startPos
}
