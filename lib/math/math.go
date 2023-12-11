package math

import (
	"math"
)

func Max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func MinElement(array []int) int {
	minVal := math.MaxInt
	for _, v := range array {
		minVal = Min(minVal, v)
	}

	return minVal
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Gcd(x, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}

func Lcm(x, b int, rest ...int) int {
	result := x * b / Gcd(x, b)

	for _, i := range rest {
		result = Lcm(result, i)
	}

	return result
}

func Sign(val int) int {
	if val == 0 {
		return 0
	}
	if val < 0 {
		return -1
	}

	return 1
}
