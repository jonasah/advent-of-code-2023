package day10

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 4, Part1(testInput1a))
	require.Equal(t, 8, Part1(testInput1b))
	require.Equal(t, 7145, Part1(realInput))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 4, Part2(testInput2a))
	require.Equal(t, 4, Part2(testInput2b))
	require.Equal(t, 8, Part2(testInput2c))
	require.Equal(t, 10, Part2(testInput2d))
	require.Equal(t, 445, Part2(realInput))
}

var realInput = common.GetInput(10)

const testInput1a = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`

const testInput1b = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

const testInput2a = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const testInput2b = `..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........`

const testInput2c = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

const testInput2d = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`
