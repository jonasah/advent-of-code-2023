package day12

import (
	"runtime"
	"strings"
	"sync"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	mathlib "github.com/jonasah/advent-of-code-2023/lib/math"
	"github.com/jonasah/advent-of-code-2023/lib/sliceconv"
)

const (
	operational = "."
	damaged     = "#"
	unknown     = "?"
)

func Part1(input string) int {
	lines := common.GetLines(input)

	sum := 0
	for _, line := range lines {
		x := strings.Split(line, " ")
		record := x[0]
		counts := sliceconv.Atoi(strings.Split(x[1], ","))

		sum += countArrangements(record, counts)
	}

	return sum
}

func Part2(input string) int {
	lines := common.GetLines(input)

	numWorkers := mathlib.Min(runtime.NumCPU()-1, len(lines))
	resChan := make(chan int, len(lines))
	workChan := make(chan int, len(lines))

	for r := range lines {
		workChan <- r
	}

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(i int) {
		loop:
			for {
				select {
				case r := <-workChan:
					x := strings.Split(lines[r], " ")
					record := repeat(x[0], unknown, 5)
					counts := sliceconv.Atoi(strings.Split(repeat(x[1], ",", 5), ","))

					res := countArrangements(record, counts)
					resChan <- res
				default:
					break loop
				}
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	res := 0
	for len(resChan) > 0 {
		res += <-resChan
	}

	return res
}

func repeat(s, delimiter string, n int) string {
	var sb strings.Builder
	sb.WriteString(s)

	for n > 1 {
		n--

		sb.WriteString(delimiter)
		sb.WriteString(s)
	}

	return sb.String()
}

func countArrangements(record string, counts []int) int {
	if len(counts) == 0 {
		if strings.ContainsAny(record, damaged) {
			return 0
		}

		return 1
	}

	idx := len(counts) / 2
	count := counts[idx]

	leftCounts := counts[:idx]
	leftSpace := minSpringsRequired(leftCounts)
	if len(leftCounts) > 0 {
		leftSpace++
	}

	rightCounts := counts[idx+1:]
	rightSpace := minSpringsRequired(rightCounts)
	if len(rightCounts) > 0 {
		rightSpace++
	}

	c := 0

	for i := leftSpace; i < len(record)-rightSpace-(count-1); i++ {
		i0 := i
		i1 := i + count
		solution := strings.Repeat(damaged, count)

		if len(leftCounts) > 0 {
			i0--
			solution = operational + solution
		}
		if len(rightCounts) > 0 {
			i1++
			solution = solution + operational
		}

		if !match(record[i0:i1], solution) {
			continue
		}

		c += countArrangements(record[:i0], leftCounts) * countArrangements(record[i1:], rightCounts)
	}

	return c
}

func minSpringsRequired(counts []int) int {
	if len(counts) == 0 {
		return 0
	}
	return mathlib.Sum(counts) + len(counts) - 1
}

func match(record, solution string) bool {
	for i, c := range record {
		if c == '?' {
			continue
		}

		if c != rune(solution[i]) {
			return false
		}
	}

	return true
}
