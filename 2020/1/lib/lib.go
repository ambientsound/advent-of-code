package lib

import (
	"strconv"
	"strings"
)

func Readints(raw []byte) []int {
	lines := strings.Split(string(raw), "\n")
	ints := make([]int, 0, len(lines))
	for _, line := range lines {
		if len(line) > 0 {
			i, _ := strconv.Atoi(line)
			ints = append(ints, i)
		}
	}
	return ints
}

func Multiply(ints []int) int {
	product := ints[0] * ints[1]
	if len(ints) > 2 {
		return Multiply(append(ints[2:], product))
	}
	return product
}

func sumslice(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

func Sums2(sum, count int, base, remainder []int) ([]int, bool) {
	if len(base)+len(remainder) < count {
		return nil, false
	}
	for i := range remainder[1:] {
		if len(base) < count-1 {
			if i+1 >= len(remainder) {
				return nil, false
			}
			result, found := Sums2(sum, count, append(base, remainder[i]), remainder[i+1:])
			if found {
				return result, found
			}
		} else if sumslice(append(base, remainder[i])) == sum {
			return append(base, remainder[i]), true
		}
	}
	return nil, false
}
