package day23

import (
	"maps"
	"slices"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
)

const (
	path       = '.'
	forest     = '#'
	slopeRight = '>'
	slopeDown  = 'v'
	slopeLeft  = '<'
	slopeUp    = '^'
)

var (
	right = Direction{1, 0}
	down  = Direction{0, 1}
	left  = Direction{-1, 0}
	up    = Direction{0, -1}

	allDirections = []Direction{right, down, left, up}

	slopeMap = map[rune]Direction{
		slopeRight: right,
		slopeDown:  down,
		slopeLeft:  left,
		slopeUp:    up,
	}
)

type Direction struct {
	x, y int
}

func (d Direction) isOpposite(other Direction) bool {
	return d.x+other.x == 0 && d.y+other.y == 0
}

type Position struct {
	x, y int
}

func (p Position) add(dir Direction) Position {
	return Position{p.x + dir.x, p.y + dir.y}
}

type TrailMap struct {
	data          [][]rune
	width, height int
}

func (t TrailMap) at(p Position) rune {
	return t.data[p.y][p.x]
}

func (t TrailMap) isValid(p Position) bool {
	return p.x >= 0 && p.x < t.width && p.y >= 0 && p.y < t.height
}

func Part1(input string) int {
	trailMap := parseInput(input)

	startPos := Position{1, 0}
	endPos := Position{trailMap.width - 2, trailMap.height - 1}

	type Hike struct {
		pos    Position
		dir    Direction
		length int
	}

	hikes := []Hike{{startPos, down, 0}}
	longestHike := 0

	for len(hikes) > 0 {
		hike := hikes[0]
		hikes = hikes[1:]

		if hike.pos == endPos {
			longestHike = math.Max(hike.length, longestHike)
			continue
		}

		for _, dir := range allDirections {
			if dir.isOpposite(hike.dir) {
				// don't go back the same way
				continue
			}

			pos := hike.pos.add(dir)

			if trailMap.at(pos) == forest {
				continue
			}

			if trailMap.at(pos) == path {
				hikes = append(hikes, Hike{pos, dir, hike.length + 1})
				continue
			}

			if dir.isOpposite(slopeMap[trailMap.at(pos)]) {
				continue
			}

			hikes = append(hikes, Hike{pos, dir, hike.length + 1})
		}
	}

	return longestHike
}

func Part2(input string) int {
	trailMap := parseInput(input)

	startPos := Position{1, 0}
	endPos := Position{trailMap.width - 2, trailMap.height - 1}

	distances := make(map[Position]map[Position]int)
	totalDistance := 0

	{
		type Hike struct {
			from   Position
			pos    Position
			dir    Direction
			length int
		}

		hikes := []Hike{{startPos, startPos, down, 0}}
		for len(hikes) > 0 {
			hike := hikes[0]
			hikes = hikes[1:]

			if hike.pos == endPos || isJunction(trailMap, hike.pos) {
				if _, ok := distances[hike.from][hike.pos]; ok {
					continue
				}

				if _, ok := distances[hike.from]; !ok {
					distances[hike.from] = make(map[Position]int)
				}
				if _, ok := distances[hike.pos]; !ok {
					distances[hike.pos] = make(map[Position]int)
				}

				distances[hike.from][hike.pos] = hike.length
				distances[hike.pos][hike.from] = hike.length
				totalDistance += hike.length

				if hike.pos != endPos {
					for _, dir := range allDirections {
						if dir.isOpposite(hike.dir) {
							// don't go back the same way
							continue
						}

						pos := hike.pos.add(dir)

						if trailMap.at(pos) == forest {
							continue
						}

						hikes = append(hikes, Hike{hike.pos, pos, dir, 1})
					}
				}

				continue
			}

			for _, dir := range allDirections {
				if dir.isOpposite(hike.dir) {
					// don't go back the same way
					continue
				}

				pos := hike.pos.add(dir)

				if trailMap.at(pos) == forest {
					continue
				}

				hikes = append(hikes, Hike{hike.from, pos, dir, hike.length + 1})
				break
			}
		}
	}

	avgDistance := totalDistance / len(distances)
	longestHike := 0

	{
		type Hike struct {
			pos     Position
			length  int
			visited map[Position]bool
		}

		hikes := []Hike{{startPos, 0, map[Position]bool{startPos: true}}}

		for len(hikes) > 0 {
			hike := hikes[0]
			hikes = hikes[1:]

			if hike.pos == endPos {
				longestHike = math.Max(longestHike, hike.length)
				continue
			}

			maxLenLimit := hike.length + (len(distances)-len(hike.visited))*avgDistance
			if maxLenLimit < longestHike {
				continue
			}

			for p, d := range distances[hike.pos] {
				if hike.visited[p] {
					continue
				}

				visited := maps.Clone(hike.visited)
				visited[p] = true
				hikes = append(hikes, Hike{p, hike.length + d, visited})
			}

			slices.SortFunc(hikes, func(a, b Hike) int { return b.length - a.length })
		}
	}

	return longestHike
}

func parseInput(input string) TrailMap {
	lines := common.GetLines(input)

	trailMap := TrailMap{width: len(lines[0]), height: len(lines)}
	for _, line := range lines {
		trailMap.data = append(trailMap.data, []rune(line))
	}

	return trailMap
}

func isJunction(trailMap TrailMap, pos Position) bool {
	if trailMap.at(pos) != path {
		return false
	}

	numNeighbors := 0
	for _, dir := range allDirections {
		p := pos.add(dir)
		if trailMap.isValid(p) && trailMap.at(p) != forest {
			numNeighbors++
		}
	}

	return numNeighbors > 2
}
