package main

import (
	"aoc2025/utils"
	"fmt"
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
	var maxX, maxY int

	var points []point
	for _, line := range lines {
		var p point
		fmt.Sscanf(line, "%d,%d", &p.y, &p.x)
		points = append(points, p)

		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)
	}

	// ok were going to just make a grid and simulate this
	grid := make([][]byte, maxX+1)
	for i := range grid {
		grid[i] = make([]byte, maxY+1)
	}

	// connect adjacent points in list
	connect(points[0], points[len(points)-1], grid)
	for i := 1; i < len(points); i++ {
		lp := points[i-1]
		p := points[i]
		connect(p, lp, grid)
	}

	// fill outer grid
	fill(grid, 0, 0)

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

func fill(grid [][]byte, i, j int) {
	if i == len(grid) || j == len(grid[0]) || grid[i][j] != 0 {
		return
	}

	grid[i][j] = '.'
	fill(grid, i+1, j)
	fill(grid, i, j+1)
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
	minX := min(p.x, q.x)
	minY := min(p.y, q.y)
	maxX := max(p.x, q.x)
	maxY := max(p.y, q.y)
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if grid[i][j] == '.' {
				return false
			}
		}
	}
	return true
}
