package day7

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 6440, Part1(testInput))
	assert.Equal(t, 250232501, Part1(realInput))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 5905, Part2(testInput))
	assert.Equal(t, 249138943, Part2(realInput))
}

var realInput = common.GetInput(7)

const testInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
