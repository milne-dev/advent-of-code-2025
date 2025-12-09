package main

import (
	"aoc2025/utils"
	"fmt"
	"math"
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
	minX := math.MaxInt
	var maxX, maxY int

	var points []point
	for _, line := range lines {
		var p point
		fmt.Sscanf(line, "%d,%d", &p.y, &p.x)
		points = append(points, p)

		minX = min(minX, p.x)

		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)
	}

	// ok were going to just make a grid and simulate this
	grid := make([][]byte, maxX+1)
	for i := range grid {
		grid[i] = make([]byte, maxY+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// connect adjacent points in list
	connect(points[0], points[len(points)-1], grid)
	for i := 1; i < len(points); i++ {
		lp := points[i-1]
		p := points[i]

		connect(p, lp, grid)
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}

	var ans int
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			q := points[j]
			if isValid(p, q, grid) {
				// l * w
				w := abs(p.x-q.x) + 1
				l := abs(p.y-q.y) + 1
				area := l * w
				ans = max(ans, area)
			}
		}
	}
	return ans
}

func connect(p, q point, grid [][]byte) {
	if p.x == q.x {
		// same row
		for j := min(p.y, q.y); j <= max(p.y, q.y); j++ {
			grid[p.x][j] = '#'
		}
	} else {
		// same column
		for k := min(p.x, q.x); k <= max(p.x, q.x); k++ {
			grid[k][p.y] = '#'
		}
	}
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func isValid(p, q point, grid [][]byte) bool {
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

	// find the first x in start col that does not have a following x (to account for sibling lines)
	// we just need to go as far as the fin x
	for x := start.x + 1; ; x++ {
		if x >= fin.x {
			break
		}

		if grid[x][start.y] == '#' && len(grid) > x+1 && grid[x+1][start.y] != '#' {
			return false
		}
	}

	// find the first y in start row that does not have a following y (to account for sibling lines)
	for y := start.y + 1; ; y++ {
		if y >= fin.y {
			break
		}

		if grid[start.x][y] == '#' && len(grid[start.x]) > y+1 && grid[start.x][y+1] != '#' {
			return false
		}
	}

	return true
}
