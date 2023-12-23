package day18

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 62, Part1(testInput))
	require.Equal(t, 62365, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 952408144115, Part2(testInput))
	require.Equal(t, 159485361249806, Part2(realInput))
}

var realInput = common.GetInput(18)

const testInput = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`
