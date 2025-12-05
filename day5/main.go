package main

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

func main() {
	fields := strings.Fields(utils.ReadInput())
	fmt.Println(PartOne(fields))
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
