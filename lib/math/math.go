package math

import (
	m "math"
)

func Max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Pow(x, y int) int {
	return int(m.Pow(float64(x), float64(y)))
}
