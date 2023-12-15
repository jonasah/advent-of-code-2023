package day15

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 1320, Part1(testInput))
	require.Equal(t, 516657, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 145, Part2(testInput))
	require.Equal(t, 210906, Part2(realInput))
}

var realInput = common.GetInput(15)

const testInput = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
