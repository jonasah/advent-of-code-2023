package day6

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 288, Part1(testInput))
	assert.Equal(t, 800280, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 71503, Part2(testInput))
	assert.Equal(t, 45128024, Part2(realInput))
}

var realInput = common.GetInput(6)

const testInput = `Time:      7  15   30
Distance:  9  40  200`
