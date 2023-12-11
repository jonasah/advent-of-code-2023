package day11

import (
	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
)

const (
	// empty  = '.'
	galaxy = '#'
)

type Position struct {
	r, c int
}

type Image struct {
	galaxies  []Position
	emptyRows []int
	emptyCols []int
}

func Solve(input string, expansion int) int {
	image := parseImage(input)

	sum := 0
	for i, p0 := range image.galaxies {
		for _, p1 := range image.galaxies[i+1:] {
			sum += manhattanDistance(p0, p1)

			for _, r := range image.emptyRows {
				if math.Sign(p0.r-r) != math.Sign(p1.r-r) {
					sum += expansion - 1
				}
			}

			for _, c := range image.emptyCols {
				if math.Sign(p0.c-c) != math.Sign(p1.c-c) {
					sum += expansion - 1
				}
			}
		}
	}

	return sum
}

func parseImage(input string) Image {
	lines := common.GetLines(input)

	var image Image

	for r, row := range lines {
		isEmptyRow := true

		for c, v := range row {
			if v == galaxy {
				image.galaxies = append(image.galaxies, Position{r, c})
				isEmptyRow = false
			}
		}

		if isEmptyRow {
			image.emptyRows = append(image.emptyRows, r)
		}
	}

	for c := range lines[0] {
		isEmptyCol := true

		for r := range lines {
			if lines[r][c] == galaxy {
				isEmptyCol = false
				break
			}
		}

		if isEmptyCol {
			image.emptyCols = append(image.emptyCols, c)
		}
	}

	return image
}

func manhattanDistance(p0, p1 Position) int {
	return math.Abs(p0.r-p1.r) + math.Abs(p0.c-p1.c)
}
