package day24

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 2, Part1(testInput, 7, 27))
	require.Equal(t, 28174, Part1(realInput, 200000000000000, 400000000000000))
}

func TestPart2(t *testing.T) {
	require.Equal(t, -1, Part2(testInput))
	// require.Equal(t, -1, Part2(realInput))
}

var realInput = common.GetInput(24)

const testInput = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`
