package day1

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 142, Part1(testInput))
	assert.Equal(t, 55477, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 281, Part2(testInput2))
	assert.Equal(t, 54431, Part2(realInput))
}

var realInput = common.GetInput(1)

const testInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const testInput2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
