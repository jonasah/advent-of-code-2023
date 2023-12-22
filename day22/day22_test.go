package day22

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 5, Part1(testInput))
	require.Equal(t, 375, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 7, Part2(testInput))
	require.Equal(t, 72352, Part2(realInput))
}

var realInput = common.GetInput(22)

const testInput = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`
