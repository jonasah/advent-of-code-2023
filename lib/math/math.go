package math

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
