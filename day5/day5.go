package day5

import (
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/jonasah/advent-of-code-2023/lib/math"
	"github.com/jonasah/advent-of-code-2023/lib/sliceconv"
)

type MapEntry struct {
	dst, src, len int
}

func (e MapEntry) containsSource(v int) bool {
	return e.src <= v && v <= e.src+e.len-1
}

func (e MapEntry) mapToDestination(v int) int {
	// assumes that v is in source range
	return e.dst + v - e.src
}

type Map struct {
	to      string
	entries []MapEntry
}

type Range struct {
	start, len int
}

func (r Range) contains(v int) bool {
	return r.start <= v && v <= r.end()
}

func (r Range) end() int {
	return r.start + r.len - 1
}

type Almanac struct {
	seeds      []int          // part 1 only
	seedRanges []Range        // part 2 only
	maps       map[string]Map // part 1: top to bottom, part 2: bottom to top
}

const (
	startCategory = "seed"
	endCategory   = "location"
)

func Part1(input string) int {
	almanac := parseInput(input, 1)

	var numbers []int
	numbers = append(numbers, almanac.seeds...)

	category := startCategory

	for category != endCategory {
		m := almanac.maps[category]
		mapped := make([]int, 0, len(numbers))

	outer:
		for _, n := range numbers {
			for _, e := range m.entries {
				if e.containsSource(n) {
					mapped = append(mapped, e.mapToDestination(n))
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
	almanac := parseInput(input, 2)

	location := 0

	for {
		category := endCategory
		number := location

		for category != startCategory {
			m := almanac.maps[category]

			for _, e := range m.entries {
				if e.containsSource(number) {
					number = e.mapToDestination(number)
					break
				}
			}

			category = m.to
		}

		for _, r := range almanac.seedRanges {
			if r.contains(number) {
				return location
			}
		}

		location++
	}
}

func parseInput(input string, part int) Almanac {
	groups := strings.Split(input, "\n\n")

	almanac := Almanac{maps: make(map[string]Map)}

	if part == 1 {
		almanac.seeds = append(almanac.seeds, sliceconv.Atoi(strings.Fields(groups[0])[1:])...)
	} else {
		ranges := sliceconv.Atoi(strings.Fields(groups[0])[1:])
		for i := 0; i < len(ranges); i += 2 {
			start := ranges[i]
			l := ranges[i+1]

			almanac.seedRanges = append(almanac.seedRanges, Range{start: start, len: l})
		}
	}

	for _, g := range groups[1:] {
		lines := common.GetLines(g)
		c := strings.Split(lines[0][:len(lines[0])-5], "-")

		from := c[0]
		to := c[2]
		if part == 2 {
			from, to = to, from
		}

		m := Map{to: to}

		for _, line := range lines[1:] {
			values := sliceconv.Atoi(strings.Fields(line))

			e := MapEntry{src: values[1], dst: values[0], len: values[2]}
			if part == 2 {
				e.src, e.dst = e.dst, e.src
			}

			m.entries = append(m.entries, e)
		}

		almanac.maps[from] = m
	}

	return almanac
}
