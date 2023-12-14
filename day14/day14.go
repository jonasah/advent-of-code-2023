package day14

import (
	"slices"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
)

const (
	empty     = "."
	cubeRock  = "#"
	roundRock = "O"
)

type Platform struct {
	data          [][]rune
	width, height int
}

func (p *Platform) row(y int) string {
	return string(p.data[y])
}

func (p *Platform) col(x int) string {
	var sb strings.Builder
	sb.Grow(p.height)
	for _, row := range p.data {
		sb.WriteRune(row[x])
	}

	return sb.String()
}

func (p *Platform) setRow(y int, values string) {
	p.data[y] = []rune(values)
}

func (p *Platform) setCol(x int, values string) {
	for y, v := range values {
		p.data[y][x] = v
	}
}

func Part1(input string) int {
	platform := parseInput(input)

	rollNorth(&platform)

	return calcLoadNorth(platform)
}

func Part2(input string) int {
	platform := parseInput(input)

	windowSize := 5
	loadHistory := make([]int, 0, 2*windowSize)

	loadCycleLength := -1
	numCycles := 1000000000
	cycle := 1

	for ; cycle <= numCycles && loadCycleLength == -1; cycle++ {
		rollNorth(&platform)
		rollWest(&platform)
		rollSouth(&platform)
		rollEast(&platform)

		loadHistory = append(loadHistory, calcLoadNorth(platform))

		// detect cycle in load history
		if len(loadHistory) < 2*windowSize {
			continue
		}

		reference := loadHistory[len(loadHistory)-windowSize:]

		for i := 0; i < len(loadHistory)-2*windowSize; i++ {
			window := loadHistory[i : i+windowSize]

			if slices.Equal(reference, window) {
				loadCycleLength = len(loadHistory) - windowSize - i
				break
			}
		}
	}

	cycle += (numCycles - cycle) / loadCycleLength * loadCycleLength

	for ; cycle <= numCycles; cycle++ {
		rollNorth(&platform)
		rollWest(&platform)
		rollSouth(&platform)
		rollEast(&platform)
	}

	return calcLoadNorth(platform)
}

func parseInput(input string) Platform {
	lines := common.GetLines(input)

	platform := Platform{width: len(lines[0]), height: len(lines)}
	for _, line := range lines {
		platform.data = append(platform.data, []rune(line))
	}

	return platform
}

func rollNorth(platform *Platform) {
	roll(platform.width, platform.col, platform.setCol, false)
}

func rollWest(platform *Platform) {
	roll(platform.height, platform.row, platform.setRow, false)
}

func rollSouth(platform *Platform) {
	roll(platform.width, platform.col, platform.setCol, true)
}

func rollEast(platform *Platform) {
	roll(platform.height, platform.row, platform.setRow, true)
}

func roll(size int, get func(i int) string, set func(i int, values string), reverse bool) {
	for i := 0; i < size; i++ {
		seq := get(i)

		sections := strings.Split(seq, cubeRock)
		for i, sec := range sections {
			numRoundRock := strings.Count(sec, roundRock)
			if reverse {
				sections[i] = strings.Repeat(empty, len(sec)-numRoundRock) + strings.Repeat(roundRock, numRoundRock)
			} else {
				sections[i] = strings.Repeat(roundRock, numRoundRock) + strings.Repeat(empty, len(sec)-numRoundRock)
			}
		}

		set(i, strings.Join(sections, cubeRock))
	}
}

func calcLoadNorth(platform Platform) int {
	load := 0
	for r, row := range platform.data {
		numRoundRocks := strings.Count(string(row), roundRock)
		load += (platform.height - r) * numRoundRocks
	}

	return load
}
