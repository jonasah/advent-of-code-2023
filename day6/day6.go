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
		count := 0
		for holdTime := 0; holdTime <= t; holdTime++ {
			speed := holdTime
			d := (t - holdTime) * speed
			if d > distances[i] {
				count++
			}
		}

		res *= count
	}

	return res
}

func Part2(input string) int {
	lines := common.GetLines(input)

	time, _ := strconv.Atoi(strings.ReplaceAll(lines[0][5:], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(lines[1][9:], " ", ""))

	minHoldTime := 0
	for holdTime := 0; holdTime <= time; holdTime++ {
		speed := holdTime
		d := (time - holdTime) * speed
		if d > distance {
			minHoldTime = holdTime
			break
		}
	}

	maxHoldTime := 0
	for holdTime := time; holdTime >= 0; holdTime-- {
		speed := holdTime
		d := (time - holdTime) * speed
		if d > distance {
			maxHoldTime = holdTime
			break
		}
	}

	return maxHoldTime - minHoldTime + 1
}
