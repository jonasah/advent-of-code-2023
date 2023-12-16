package day16

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 46, Part1(testInput))
	require.Equal(t, 7498, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 51, Part2(testInput))
	require.Equal(t, 7846, Part2(realInput))
}

var realInput = common.GetInput(16)

const testInput = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`
