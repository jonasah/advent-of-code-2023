package day24

import (
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2023/lib/common"
)

type Vector struct {
	x, y float64
}

func (v Vector) dot(other Vector) float64 {
	return v.x*other.x + v.y*other.y
}

func (v Vector) add(other Vector) Vector {
	return Vector{v.x + other.x, v.y + other.y}
}

func (v Vector) minus(other Vector) Vector {
	return Vector{v.x - other.x, v.y - other.y}
}

func (v Vector) mult(d float64) Vector {
	return Vector{v.x * d, v.y * d}
}

type Hailstone struct {
	x, y, z    float64
	vx, vy, vz float64
}

func Part1(input string, min, max int) int {
	hailstones := parseInput(input)

	minf := float64(min)
	maxf := float64(max)

	sum := 0
	for i, h1 := range hailstones {
		for _, h2 := range hailstones[i+1:] {
			p1 := Vector{h1.x, h1.y}
			v1 := Vector{h1.vx, h1.vy}
			n1 := Vector{v1.y, -v1.x}

			p2 := Vector{h2.x, h2.y}
			v2 := Vector{h2.vx, h2.vy}
			n2 := Vector{v2.y, -v2.x}

			t1 := float64(n2.dot(p2.minus(p1))) / float64(n2.dot(v1))
			t2 := float64(n1.dot(p1.minus(p2))) / float64(n1.dot(v2))

			pi := p1.add(v1.mult(t1))

			if t1 >= 0 && t2 > 0 && pi.x >= minf && pi.x <= maxf && pi.y >= minf && pi.y <= maxf {
				sum++
			}
		}
	}

	return sum
}

func Part2(input string) int {
	return 0
}

func parseInput(input string) []Hailstone {
	lines := common.GetLines(input)

	var hailstones []Hailstone
	for _, line := range lines {
		fields := strings.FieldsFunc(line, func(r rune) bool { return strings.ContainsAny(string(r), ", @") })
		x, _ := strconv.ParseFloat(fields[0], 64)
		y, _ := strconv.ParseFloat(fields[1], 64)
		z, _ := strconv.ParseFloat(fields[2], 64)
		vx, _ := strconv.ParseFloat(fields[3], 64)
		vy, _ := strconv.ParseFloat(fields[4], 64)
		vz, _ := strconv.ParseFloat(fields[5], 64)

		hailstones = append(hailstones, Hailstone{x, y, z, vx, vy, vz})
	}

	return hailstones
}
