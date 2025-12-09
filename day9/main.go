package main

import (
	"aoc2025/utils"
	"fmt"
	"slices"
)

func main() {
	input := utils.ReadInput()
	lines := utils.StringLines(input)
	fmt.Println(PartTwo(lines))
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
	// farthest red to left for row
	prefix := make(map[int]int)
	suffix := make(map[int]int)

	var maxX, maxY int

	var points []point
	for _, line := range lines {
		var p point
		fmt.Sscanf(line, "%d,%d", &p.y, &p.x)
		points = append(points, p)

		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)

		if m, ok := prefix[p.x]; ok {
			prefix[p.x] = min(m, p.y)
		} else {
			prefix[p.x] = p.y
		}

		if m, ok := suffix[p.x]; ok {
			suffix[p.x] = max(m, p.y)
		} else {
			suffix[p.x] = p.y
		}
	}

	// ok were going to just make a grid and simulate this
	grid := make([][]byte, maxX+1)
	for i := range grid {
		grid[i] = make([]byte, maxY+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// lets just map this out a bit
	// for x []y
	horizPoints := make(map[int][]int)
	// for y []x
	vertPoints := make(map[int][]int)
	for _, p := range points {
		horizPoints[p.x] = append(horizPoints[p.x], p.y)
		vertPoints[p.y] = append(vertPoints[p.y], p.x)
	}

	// now lets color in the lines
	for i, js := range horizPoints {
		minJ := slices.Min(js)
		maxJ := slices.Max(js)
		for j := minJ; j <= maxJ; j++ {
			grid[i][j] = '#'
		}
	}

	for j, is := range vertPoints {
		minI := slices.Min(is)
		maxI := slices.Max(is)
		for i := minI; i <= maxI; i++ {
			grid[i][j] = '#'
		}
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}

	// ok time to capture min and max of each line

	var ans int
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			q := points[j]
			// l * w
			w := max(p.y, q.y) - min(p.y, q.y) + 1
			l := max(p.x, q.x) - min(p.x, q.x) + 1
			area := l * w
			if isValid(p, q, prefix, suffix) {
				ans = max(ans, area)
			}
		}
	}
	return ans
}

func isValid(p, q point, prefix, suffix map[int]int) bool {
	var start, fin point

	if p.x == q.x {
		if p.y < q.y {
			start = p
			fin = q
		} else {
			start = q
			fin = p
		}
	} else if p.x < q.x {
		start = p
		fin = q
	} else {
		start = q
		fin = p
	}

	for x := start.x; x <= fin.x; x++ {
		if pref, ok := prefix[x]; ok && pref > start.x {
			return false
		}
		if suf, ok := suffix[x]; ok && suf < fin.y {
			return false
		}
	}

	return true
}
