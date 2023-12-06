package day6

import (
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/sliceconv"
)

func Part1(input string) int {
	lines := common.GetLines(input)

	times := sliceconv.Atoi(strings.Fields(lines[0][5:]))
	distances := sliceconv.Atoi(strings.Fields(lines[1][9:]))

	res := 1
	for i, t := range times {
		res *= getNumberOfWaysToWin(t, distances[i])
	}

	return res
}

func Part2(input string) int {
	lines := common.GetLines(input)

	time, _ := strconv.Atoi(strings.ReplaceAll(lines[0][5:], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(lines[1][9:], " ", ""))

	return getNumberOfWaysToWin(time, distance)
}

func getNumberOfWaysToWin(time, distance int) int {
	for holdTime := 0; holdTime <= time; holdTime++ {
		speed := holdTime
		d := (time - holdTime) * speed
		if d > distance {
			return (time - holdTime) - holdTime + 1
		}
	}

	return 0
}
