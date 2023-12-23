package day18

import (
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
)

type Direction struct {
	x, y int
}

func (d Direction) scale(n int) Direction {
	return Direction{d.x * n, d.y * n}
}

type Position struct {
	x, y int
}

func (p Position) add(dir Direction) Position {
	return Position{p.x + dir.x, p.y + dir.y}
}

var (
	right = Direction{1, 0}
	down  = Direction{0, 1}
	left  = Direction{-1, 0}
	up    = Direction{0, -1}

	dirMap = map[byte]Direction{
		// part 1
		'R': right,
		'D': down,
		'L': left,
		'U': up,
		// part 2
		'0': right,
		'1': down,
		'2': left,
		'3': up,
	}
)

func Part1(input string) int {
	return calculateTrenchSize(input, 1)
}

func Part2(input string) int {
	return calculateTrenchSize(input, 2)
}

func calculateTrenchSize(input string, part int) int {
	points, edgeLength := getPoints(input, part)

	sum := 0
	for i, p0 := range points[:len(points)-1] {
		p1 := points[i+1]
		sum += p0.x*p1.y - p1.x*p0.y
	}

	return (sum + edgeLength + 2) / 2
}

func getPoints(input string, part int) ([]Position, int) {
	lines := common.GetLines(input)

	points := []Position{{0, 0}}
	edgeLength := 0
	for _, line := range lines {
		s := strings.FieldsFunc(line, func(r rune) bool { return strings.ContainsAny(string(r), " ()#") })

		dir := dirMap[s[0][0]]
		dist, _ := strconv.ParseInt(s[1], 10, 0)
		color := s[2]

		if part == 2 {
			dir = dirMap[color[len(color)-1]]
			dist, _ = strconv.ParseInt(color[:len(color)-1], 16, 0)
		}

		points = append(points, points[len(points)-1].add(dir.scale(int(dist))))
		edgeLength += int(dist)
	}

	return points, edgeLength
}
