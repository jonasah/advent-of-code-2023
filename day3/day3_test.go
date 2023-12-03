package day3

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 4361, Part1(testInput))
	assert.Equal(t, 509115, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 467835, Part2(testInput))
	assert.Equal(t, 75220503, Part2(realInput))
}

var realInput = common.GetInput(3)

const testInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
