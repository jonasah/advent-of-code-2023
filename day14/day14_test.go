package day14

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 136, Part1(testInput))
	require.Equal(t, 110565, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 64, Part2(testInput))
	require.Equal(t, 89845, Part2(realInput))
}

var realInput = common.GetInput(14)

const testInput = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`
