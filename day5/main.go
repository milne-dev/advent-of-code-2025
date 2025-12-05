package main

import (
	"aoc2025/utils"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	fields := strings.Fields(utils.ReadInput())
	fmt.Println(PartTwo(fields))
}

func PartOne(fields []string) int {
	ranges := make([][2]int, 0, len(fields))
	var ans int

	for _, line := range fields {
		var n1, n2 int
		var onN2 bool
		for _, c := range line {
			if c == '-' {
				onN2 = true
				continue
			}

			if onN2 {
				n2 = n2*10 + int(c-'0')
			} else {
				n1 = n1*10 + int(c-'0')
			}
		}

		if onN2 {
			ranges = append(ranges, [2]int{n1, n2})
		} else if isWithinRanges(ranges, n1) {
			ans++
		}
	}

	return ans
}

func isWithinRanges(ranges [][2]int, n int) bool {
	for _, r := range ranges {
		if n >= r[0] && n <= r[1] {
			return true
		}
	}
	return false
}

func PartTwo(fields []string) int {
	ranges := make([][2]int, 0, len(fields))

	for _, line := range fields {
		var n1, n2 int
		var onN2 bool
		for _, c := range line {
			if c == '-' {
				onN2 = true
				continue
			}

			if onN2 {
				n2 = n2*10 + int(c-'0')
			} else {
				n1 = n1*10 + int(c-'0')
			}
		}

		ranges = append(ranges, [2]int{n1, n2})
	}

	// combine ranges
	slices.SortFunc(ranges, func(a, b [2]int) int {
		x := cmp.Compare(a[0], b[0])
		if x == 0 {
			return cmp.Compare(a[1], b[1])
		}
		return x
	})

	merged := make([][2]int, 0, len(ranges))
	for i, r := range ranges {
		start, end := r[0], r[1]

		if len(merged) != 0 && merged[len(merged)-1][1] >= end {
			continue
		}

		for j := i + 1; j < len(ranges); j++ {
			if ranges[j][0] <= end {
				end = max(ranges[j][1], end)
			}
		}

		merged = append(merged, [2]int{start, end})
	}

	var ans int
	for _, r := range merged {
		ans += r[1] - r[0] + 1
	}
	return ans
}
