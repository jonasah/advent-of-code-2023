package day11

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 374, Solve(testInput, 2))
	require.Equal(t, 9563821, Solve(realInput, 2))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 1030, Solve(testInput, 10))
	require.Equal(t, 8410, Solve(testInput, 100))
	require.Equal(t, 827009909817, Solve(realInput, 1000000))
}

var realInput = common.GetInput(11)

const testInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
