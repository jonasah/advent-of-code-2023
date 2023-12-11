package day10

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/color"
	"github.com/jonasah/advent-of-code-2023/lib/common"
)

const (
	vertical   = '|'
	horizontal = '-'

	eastToSouth = '7'
	eastToNorth = 'J'

	southToEast = 'L'
	southToWest = 'J'

	westToSouth = 'F'
	westToNorth = 'L'

	northToEast = 'F'
	northToWest = '7'

	ground = '.'
	start  = 'S'
)

type Direction int

const (
	east Direction = iota
	south
	west
	north
)

type Position struct {
	x, y int
}

type Tile struct {
	pipe rune

	isPadding bool
	isInLoop  bool
	isOutside bool
}

type Grid struct {
	width, height int
	tiles         [][]Tile
}

func NewGrid(w, h int) Grid {
	grid := Grid{width: w, height: h}

	for y := 0; y < h; y++ {
		row := make([]Tile, 0, w)
		for x := 0; x < w; x++ {
			row = append(row, Tile{pipe: ground, isPadding: true})
		}

		grid.tiles = append(grid.tiles, row)
	}

	return grid
}

func (g Grid) at(p Position) Tile {
	return g.tiles[p.y][p.x]
}

func (g Grid) isValid(p Position) bool {
	return p.x >= 0 && p.x < g.width && p.y >= 0 && p.y < g.height
}

func (g Grid) setPipe(p Position, pipe rune, isPadding bool) {
	g.tiles[p.y][p.x].pipe = pipe
	g.tiles[p.y][p.x].isPadding = isPadding
}

func (g Grid) markInLoop(p Position) {
	g.tiles[p.y][p.x].isInLoop = true
}

func (g Grid) markOutside(p Position) {
	g.tiles[p.y][p.x].isOutside = true
}

func Part1(input string) int {
	grid := parseGrid(input)
	return findLoop(&grid) / 3 / 2
}

func Part2(input string) int {
	grid := parseGrid(input)
	findLoop(&grid)

	toVisit := []Position{{0, 0}}

	for len(toVisit) > 0 {
		p := toVisit[0]
		toVisit = toVisit[1:]

		if !grid.isValid(p) {
			continue
		}

		tile := grid.at(p)
		if tile.isOutside || tile.isInLoop {
			continue
		}

		grid.markOutside(p)

		toVisit = append(
			toVisit,
			Position{p.x - 1, p.y},
			Position{p.x + 1, p.y},
			Position{p.x, p.y - 1},
			Position{p.x, p.y + 1})
	}

	numInside := 0
	for _, row := range grid.tiles {
		for _, tile := range row {
			if !tile.isInLoop && !tile.isOutside && !tile.isPadding {
				numInside++
			}
		}
	}

	// dump(grid)

	return numInside
}

func parseGrid(input string) Grid {
	lines := common.GetLines(input)

	w := 3 * len(lines[0])
	h := 3 * len(lines)
	grid := NewGrid(w, h)

	for y, line := range lines {
		for x, pipe := range line {
			center := Position{3*x + 1, 3*y + 1}
			grid.setPipe(center, pipe, false)

			left := Position{center.x - 1, center.y}
			right := Position{center.x + 1, center.y}
			top := Position{center.x, center.y - 1}
			bottom := Position{center.x, center.y + 1}

			switch pipe {
			case vertical:
				grid.setPipe(top, vertical, true)
				grid.setPipe(bottom, vertical, true)
			case horizontal:
				grid.setPipe(left, horizontal, true)
				grid.setPipe(right, horizontal, true)
			case eastToSouth:
				grid.setPipe(left, horizontal, true)
				grid.setPipe(bottom, vertical, true)
			case eastToNorth:
				grid.setPipe(left, horizontal, true)
				grid.setPipe(top, vertical, true)
			case westToSouth:
				grid.setPipe(right, horizontal, true)
				grid.setPipe(bottom, vertical, true)
			case westToNorth:
				grid.setPipe(right, horizontal, true)
				grid.setPipe(top, vertical, true)
			case start:
				grid.setPipe(left, horizontal, true)
				grid.setPipe(right, horizontal, true)
				grid.setPipe(top, vertical, true)
				grid.setPipe(bottom, vertical, true)
			case ground:
			default:
				panic(fmt.Sprintf("unhandled pipe: %c", pipe))
			}
		}
	}

	for x := 0; x < w; x++ {
		grid.setPipe(Position{x, 0}, ground, true)
		grid.setPipe(Position{x, h - 1}, ground, true)
	}
	for y := 0; y < h; y++ {
		grid.setPipe(Position{0, y}, ground, true)
		grid.setPipe(Position{w - 1, y}, ground, true)
	}

	return grid
}

func findLoop(grid *Grid) int {
	startPos := findStartPos(*grid)

	pos := startPos
	initialDir := east
	dir := initialDir

	var loop []Position

	for {
		nextPos, nextDir, err := advance(*grid, pos, dir)
		if err != nil {
			pos = startPos
			initialDir++
			dir = initialDir
			loop = nil
			continue
		}

		loop = append(loop, nextPos)

		pos = nextPos
		dir = nextDir

		if pos == startPos {
			break
		}
	}

	if initialDir != east && dir != west {
		grid.setPipe(Position{startPos.x + 1, startPos.y}, ground, true)
	}
	if initialDir != west && dir != east {
		grid.setPipe(Position{startPos.x - 1, startPos.y}, ground, true)
	}
	if initialDir != north && dir != south {
		grid.setPipe(Position{startPos.x, startPos.y - 1}, ground, true)
	}
	if initialDir != south && dir != north {
		grid.setPipe(Position{startPos.x, startPos.y + 1}, ground, true)
	}

	for _, p := range loop {
		grid.markInLoop(p)
	}

	return len(loop)
}

func findStartPos(grid Grid) Position {
	for y, row := range grid.tiles {
		x := slices.IndexFunc(row, func(t Tile) bool { return t.pipe == start })
		if x != -1 {
			return Position{x, y}
		}
	}

	panic("found no start")
}

func advance(grid Grid, p Position, dir Direction) (Position, Direction, error) {
	if dir == east {
		nextPos := Position{p.x + 1, p.y}

		if !grid.isValid(nextPos) {
			return p, dir, errors.New("no path to the east")
		}

		nextPipe := grid.at(nextPos).pipe
		if nextPipe == start || nextPipe == horizontal {
			return nextPos, dir, nil
		}
		if nextPipe == eastToSouth {
			return nextPos, south, nil
		}
		if nextPipe == eastToNorth {
			return nextPos, north, nil
		}

		return p, dir, fmt.Errorf("invalid tile while heading east: %c at %v", nextPipe, nextPos)
	}

	if dir == south {
		nextPos := Position{p.x, p.y + 1}

		if !grid.isValid(nextPos) {
			return p, dir, errors.New("no path to the south")
		}

		nextPipe := grid.at(nextPos).pipe
		if nextPipe == start || nextPipe == vertical {
			return nextPos, dir, nil
		}
		if nextPipe == southToWest {
			return nextPos, west, nil
		}
		if nextPipe == southToEast {
			return nextPos, east, nil
		}

		return p, dir, fmt.Errorf("invalid tile while heading south: %c at %v", nextPipe, nextPos)
	}

	if dir == west {
		nextPos := Position{p.x - 1, p.y}

		if !grid.isValid(nextPos) {
			return p, dir, errors.New("no path to the west")
		}

		nextPipe := grid.at(nextPos).pipe
		if nextPipe == start || nextPipe == horizontal {
			return nextPos, dir, nil
		}
		if nextPipe == westToNorth {
			return nextPos, north, nil
		}
		if nextPipe == westToSouth {
			return nextPos, south, nil
		}

		return p, dir, fmt.Errorf("invalid tile while heading west: %c at %v", nextPipe, nextPos)
	}

	// north
	nextPos := Position{p.x, p.y - 1}

	if !grid.isValid(nextPos) {
		return p, dir, errors.New("no path to the north")
	}

	nextPipe := grid.at(nextPos).pipe
	if nextPipe == start || nextPipe == vertical {
		return nextPos, dir, nil
	}
	if nextPipe == northToEast {
		return nextPos, east, nil
	}
	if nextPipe == northToWest {
		return nextPos, west, nil
	}

	return p, dir, fmt.Errorf("invalid tile while heading north: %c at %v", nextPipe, nextPos)
}

func dump(grid Grid) {
	var out strings.Builder
	for y, row := range grid.tiles {
		if y%3 != 1 {
			continue
		}

		for x, tile := range row {
			if x%3 != 1 {
				continue
			}

			pipe := tile.pipe
			var colorFunc func(string) string

			if tile.isInLoop {
				colorFunc = color.Blue
			} else if tile.isOutside {
				pipe = '@'
				colorFunc = color.Green
			} else {
				pipe = '&'
				colorFunc = color.Yellow
			}

			out.WriteString(colorFunc(string(pipe)))
		}

		out.WriteRune('\n')
	}

	fmt.Println(out.String())
}
