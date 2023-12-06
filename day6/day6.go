package day6

import (
	"math"
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
	threshold := (float64(time) - math.Sqrt(float64(time*time-4*distance))) / 2
	holdTime := int(math.Ceil(threshold + 1e-9))
	return (time - holdTime) - holdTime + 1
}
