package day9

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 114, Solve(testInput, 1))
	assert.Equal(t, 1969958987, Solve(realInput, 1))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 2, Solve(testInput, 2))
	assert.Equal(t, 1068, Solve(realInput, 2))
}

var realInput = common.GetInput(9)

const testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
