package day12

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 21, Part1(testInput))
	require.Equal(t, 7792, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 525152, Part2(testInput))
	require.Equal(t, 13012052341533, Part2(realInput))
}

var realInput = common.GetInput(12)

const testInput = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`
