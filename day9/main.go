package main

import (
	"aoc2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput()
	lines := utils.StringLines(input)
	fmt.Println(PartOne(lines))
}

type point struct {
	x, y int
}

func PartOne(lines []string) int {
	var points []point
	for _, line := range lines {
		var p point
		fmt.Sscanf(line, "%d,%d", &p.x, &p.y)
		points = append(points, p)
	}

	var ans int
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			q := points[j]
			// l * w
			w := max(p.y, q.y) - min(p.y, q.y) + 1
			l := max(p.x, q.x) - min(p.x, q.x) + 1
			area := l * w
			ans = max(ans, area)
		}
	}
	return ans
}

func PartTwo(lines []string) int {
	return 0
}
