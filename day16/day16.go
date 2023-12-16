package day16

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
)

const (
	empty              = '.'
	mirrorRight        = '/'
	mirrorLeft         = '\\'
	horizontalSplitter = '-'
	verticalSplitter   = '|'

	right Direction = 1 << 0
	down  Direction = 1 << 1
	left  Direction = 1 << 2
	up    Direction = 1 << 3
)

var outgoingDirMap = map[Direction]map[rune]Direction{
	right: {
		empty:              right,
		mirrorRight:        up,
		mirrorLeft:         down,
		verticalSplitter:   up | down,
		horizontalSplitter: right,
	},
	down: {
		empty:              down,
		mirrorRight:        left,
		mirrorLeft:         right,
		verticalSplitter:   down,
		horizontalSplitter: left | right,
	},
	left: {
		empty:              left,
		mirrorRight:        down,
		mirrorLeft:         up,
		verticalSplitter:   up | down,
		horizontalSplitter: left,
	},
	up: {
		empty:              up,
		mirrorRight:        right,
		mirrorLeft:         left,
		verticalSplitter:   up,
		horizontalSplitter: left | right,
	},
}

type Direction int

func (d Direction) count() int {
	return bits.OnesCount(uint(d))
}

func (d Direction) contains(dir Direction) bool {
	return d&dir > 0
}

type Position struct {
	x, y int
}

type Tile struct {
	content rune
	beams   Direction
}

type Grid struct {
	data          [][]Tile
	width, height int
}

func (g Grid) isValid(p Position) bool {
	return p.x >= 0 && p.x < g.width && p.y >= 0 && p.y < g.height
}

func (g Grid) at(p Position) Tile {
	return g.data[p.y][p.x]
}

func (g Grid) addBeam(beam Beam) {
	g.data[beam.pos.y][beam.pos.x].beams |= beam.dir
}

func (g Grid) resetBeams() {
	for y, row := range g.data {
		for x := range row {
			g.data[y][x].beams = 0
		}
	}
}

type Beam struct {
	pos Position
	dir Direction
}

func Part1(input string) int {
	grid := parseInput(input)

	return startBeam(grid, Beam{Position{0, 0}, right})
}

func Part2(input string) int {
	grid := parseInput(input)

	var startBeams []Beam
	for x := 0; x < grid.width; x++ {
		startBeams = append(startBeams, Beam{Position{x, 0}, down}, Beam{Position{x, grid.height - 1}, up})
	}
	for y := 0; y < grid.height; y++ {
		startBeams = append(startBeams, Beam{Position{0, y}, right}, Beam{Position{grid.width - 1, y}, left})
	}

	result := 0
	for _, beam := range startBeams {
		grid.resetBeams()
		result = math.Max(result, startBeam(grid, beam))
	}

	return result
}

func parseInput(input string) Grid {
	lines := common.GetLines(input)
	grid := Grid{width: len(lines[0]), height: len(lines)}

	for _, line := range lines {
		var row []Tile
		for _, c := range line {
			row = append(row, Tile{content: c})
		}

		grid.data = append(grid.data, row)
	}

	return grid
}

func startBeam(grid Grid, beam Beam) int {
	beams := []Beam{beam}

	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]

		if !grid.isValid(beam.pos) {
			continue
		}

		tile := grid.at(beam.pos)

		if tile.beams.contains(beam.dir) {
			continue
		}

		grid.addBeam(beam)

		beams = append(beams, advance(beam, tile)...)
	}

	// dump(grid)

	sum := 0
	for _, row := range grid.data {
		for _, tile := range row {
			if tile.beams.count() > 0 {
				sum++
			}
		}
	}

	return sum
}

func advance(beam Beam, tile Tile) []Beam {
	newDir := outgoingDirMap[beam.dir][tile.content]

	var newBeams []Beam

	if newDir.contains(right) {
		newBeams = append(newBeams, Beam{Position{beam.pos.x + 1, beam.pos.y}, right})
	}
	if newDir.contains(down) {
		newBeams = append(newBeams, Beam{Position{beam.pos.x, beam.pos.y + 1}, down})
	}
	if newDir.contains(left) {
		newBeams = append(newBeams, Beam{Position{beam.pos.x - 1, beam.pos.y}, left})
	}
	if newDir.contains(up) {
		newBeams = append(newBeams, Beam{Position{beam.pos.x, beam.pos.y - 1}, up})
	}

	return newBeams
}

func dump(grid Grid) {
	dirMap := map[Direction]rune{
		right: '>',
		down:  'v',
		left:  '<',
		up:    '^',
	}

	var sb strings.Builder
	for _, row := range grid.data {
		for _, tile := range row {
			if tile.beams.count() == 0 || tile.content != empty {
				sb.WriteRune(tile.content)
				continue
			}

			if tile.beams.count() == 1 {
				sb.WriteRune(dirMap[tile.beams])
				continue
			}

			sb.WriteString(strconv.Itoa(tile.beams.count()))
		}

		sb.WriteRune('\n')
	}

	fmt.Println(sb.String())
}
