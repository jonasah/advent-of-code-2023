package day13

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 405, Part1(testInput))
	require.Equal(t, 33356, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 400, Part2(testInput))
	require.Equal(t, 28475, Part2(realInput))
}

var realInput = common.GetInput(13)

const testInput = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
