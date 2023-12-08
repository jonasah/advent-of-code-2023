package day8

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 6, Part1(testInput))
	assert.Equal(t, 21409, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 6, Part2(testInput2))
	assert.Equal(t, 21165830176709, Part2(realInput))
}

var realInput = common.GetInput(8)

const testInput = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const testInput2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
