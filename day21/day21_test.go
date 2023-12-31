package day21

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 16, Part1(testInput, 6))
	require.Equal(t, 3632, Part1(realInput, 64))
}

func TestPart2(t *testing.T) {
	require.Equal(t, -1, Part2(testInput))
	// require.Equal(t, -1, Part2(realInput))
}

var realInput = common.GetInput(21)

const testInput = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`
