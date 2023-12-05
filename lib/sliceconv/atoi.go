package sliceconv

import "strconv"

func Atoi(a []string) []int {
	l := make([]int, 0, len(a))
	for _, s := range a {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		l = append(l, i)
	}

	return l
}
