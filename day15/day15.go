package day15

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) int {
	steps := strings.Split(strings.TrimSpace(input), ",")

	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}

	return sum
}

func Part2(input string) int {
	steps := strings.Split(strings.TrimSpace(input), ",")

	type Lens struct {
		label       string
		focalLength int
	}

	hashMap := make(map[int][]Lens)

	for _, step := range steps {
		operIdx := strings.IndexAny(step, "-=")
		label := step[:operIdx]
		operation := step[operIdx]

		labelHash := hash(label)
		labelIdx := slices.IndexFunc(hashMap[labelHash], func(lens Lens) bool { return lens.label == label })

		if operation == '=' {
			focalLength, _ := strconv.Atoi(step[operIdx+1:])

			if labelIdx == -1 {
				hashMap[labelHash] = append(hashMap[labelHash], Lens{label, focalLength})
			} else {
				hashMap[labelHash][labelIdx].focalLength = focalLength
			}

			continue
		}

		// operation == '-'
		if labelIdx != -1 {
			hashMap[labelHash] = slices.Delete(hashMap[labelHash], labelIdx, labelIdx+1)
		}
	}

	power := 0
	for box, lenses := range hashMap {
		for i, lens := range lenses {
			power += (box + 1) * (i + 1) * lens.focalLength
		}
	}

	return power
}

func hash(s string) int {
	currentValue := 0
	for _, c := range s {
		currentValue = ((currentValue + int(c)) * 17) % 256
	}
	return currentValue
}
