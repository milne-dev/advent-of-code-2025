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
		p.x++
		p.y++
		points = append(points, p)

		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)
	}

	// ok were going to just make a grid and simulate this
	grid := make([][]byte, maxX+3)
	for i := range grid {
		grid[i] = make([]byte, maxY+3)
	}

	// connect adjacent points in list
	connect(points[0], points[len(points)-1], grid)
	for i := 1; i < len(points); i++ {
		lp := points[i-1]
		p := points[i]
		connect(p, lp, grid)
	}

	// fill outer grid
	fill(grid)

	// lets process the grid and find for each row that contains a point
	// mark down the location of any . that are adjacent to # or X
	// then we can check if any of those . locations fall within our range
	adjacentDotLocationsX := make([][]int, len(grid))
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '.' {
				if (j > 0 && grid[i][j-1] != '.') || (j < len(grid[i])-1 && grid[i][j+1] != '.') {
					adjacentDotLocationsX[i] = append(adjacentDotLocationsX[i], j)
				}
			}
		}
	}

	adjacentDotLocationsY := make([][]int, len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '.' {
				if (i > 0 && grid[i-1][j] != '.') || (i < len(grid)-1 && grid[i+1][j] != '.') {
					adjacentDotLocationsY[j] = append(adjacentDotLocationsY[j], i)
				}
			}
		}
	}

	var ans int
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			q := points[j]
			if isValid(p, q, adjacentDotLocationsX, adjacentDotLocationsY) {
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

func fill(grid [][]byte) {
	q := [][2]int{{0, 0}}

	for len(q) > 0 {
		var nq [][2]int

		for _, pos := range q {
			i, j := pos[0], pos[1]

			if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] != 0 {
				continue
			}

			grid[i][j] = '.'

			nq = append(nq, [2]int{i + 1, j})
			nq = append(nq, [2]int{i, j + 1})
			nq = append(nq, [2]int{i - 1, j})
			nq = append(nq, [2]int{i, j - 1})
		}

		q = nq
	}
}

func connect(p, q point, grid [][]byte) {
	grid[p.x][p.y] = '#'
	grid[q.x][q.y] = '#'

	if p.x == q.x {
		// same row
		for j := min(p.y, q.y) + 1; j < max(p.y, q.y); j++ {
			grid[p.x][j] = 'X'
		}
	} else {
		// same column
		for k := min(p.x, q.x) + 1; k < max(p.x, q.x); k++ {
			grid[k][p.y] = 'X'
		}
	}
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func isValid(p, q point, adjacentDotLocationsX, adjacentDotLocationsY [][]int) bool {
	minX := min(p.x, q.x)
	minY := min(p.y, q.y)
	maxX := max(p.x, q.x)
	maxY := max(p.y, q.y)
	for i := minX; i <= maxX; i++ {
		for _, dotJ := range adjacentDotLocationsX[i] {
			if dotJ >= minY && dotJ <= maxY {
				return false
			}
		}
	}
	for j := minY; j <= maxY; j++ {
		for _, dotI := range adjacentDotLocationsY[j] {
			if dotI >= minX && dotI <= maxX {
				return false
			}
		}
	}
	return true
}
