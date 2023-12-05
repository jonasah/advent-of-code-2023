package day5

import (
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
)

type Entry struct {
	dst int
	src int
	len int
}

func (e Entry) isInSrcRange(v int) bool {
	return e.src <= v && v <= e.src+e.len-1
}

func (e Entry) mapToDst(v int) int {
	return e.dst + v - e.src
}

type Map struct {
	from    string
	to      string
	entries []Entry
}

type Almanac struct {
	// part 1
	seeds []int
	maps  map[string]Map

	// part 2
	seedsRanges []Entry
	reverseMaps map[string]Map
}

func Part1(input string) int {
	almanac := parseInput(input, true)

	var numbers []int
	numbers = append(numbers, almanac.seeds...)

	category := "seed"

	for {
		if category == "location" {
			break
		}

		m := almanac.maps[category]
		mapped := make([]int, 0, len(numbers))

	outer:
		for _, n := range numbers {
			for _, e := range m.entries {
				if e.isInSrcRange(n) {
					mapped = append(mapped, e.mapToDst(n))
					continue outer
				}
			}

			mapped = append(mapped, n)
		}

		category = m.to
		numbers = mapped
	}

	return math.MinElement(numbers)
}

func Part2(input string) int {
	almanac := parseInput(input, false)

	category := "location"
	location := 0
	number := location

	for {
		if category == "seed" {
			for _, r := range almanac.seedsRanges {
				if r.isInSrcRange(number) {
					return location
				}
			}

			category = "location"
			location++
			number = location
		}

		m := almanac.reverseMaps[category]
		mapped := 0

		for _, e := range m.entries {
			if e.isInSrcRange(number) {
				mapped = e.mapToDst(number)
				break
			}

			mapped = number
		}

		category = m.to
		number = mapped
	}
}

func parseInput(input string, part1 bool) Almanac {
	groups := strings.Split(input, "\n\n")

	almanac := Almanac{seeds: make([]int, 0), maps: make(map[string]Map), reverseMaps: make(map[string]Map)}

	if part1 {
		almanac.seeds = append(almanac.seeds, common.ToInts(strings.Fields(groups[0])[1:])...)
	} else {
		ranges := common.ToInts(strings.Fields(groups[0])[1:])
		for i := 0; i < len(ranges); i += 2 {
			start := ranges[i]
			l := ranges[i+1]

			almanac.seedsRanges = append(almanac.seedsRanges, Entry{src: start, dst: start, len: l})
		}
	}

	for _, g := range groups[1:] {
		lines := common.GetLines(g)
		c := strings.Split(lines[0][:len(lines[0])-5], "-")

		m := Map{from: c[0], to: c[2], entries: make([]Entry, 0)}
		r := Map{from: m.to, to: m.from, entries: make([]Entry, 0)}

		for _, line := range lines[1:] {
			values := common.ToInts(strings.Fields(line))
			m.entries = append(m.entries, Entry{src: values[1], dst: values[0], len: values[2]})
			r.entries = append(r.entries, Entry{src: values[0], dst: values[1], len: values[2]})
		}

		almanac.maps[m.from] = m
		almanac.reverseMaps[r.from] = r
	}

	return almanac
}
