package main

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

var dirs = [][]int{{1, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func main() {
	grid := strings.Fields(utils.ReadInput())
	gridBytes := make([][]byte, len(grid))
	for i := range grid {
		gridBytes[i] = []byte(grid[i])
	}

	fmt.Println(PartTwo(gridBytes))
}

func PartTwo(grid [][]byte) int {
	var total int
	var removed = true

	for removed {
		removed = false
		for x := range grid {
			for y := range grid[x] {
				if grid[x][y] == '@' && canAccessP2(grid, x, y) {
					total++
					grid[x][y] = 'x'
					removed = true
				}
			}
		}
	}

	return total
}

func canAccessP2(grid [][]byte, x, y int) bool {
	var total int
	for _, dir := range dirs {
		i, j := dir[0], dir[1]
		a, b := x+i, y+j
		if a < 0 || a == len(grid) || b < 0 || b == len(grid[a]) {
			continue
		}
		if grid[a][b] == '@' {
			total++
		}
	}
	return total < 4
}

func PartOne(grid []string) int {
	var total int
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '@' && canAccess(grid, x, y) {
				total++
			}
		}
	}
	return total
}

// can access if fewer than 4 in 8 edges
func canAccess(grid []string, x, y int) bool {
	var total int
	for _, dir := range dirs {
		i, j := dir[0], dir[1]
		a, b := x+i, y+j
		if a < 0 || a == len(grid) || b < 0 || b == len(grid[a]) {
			continue
		}
		if grid[a][b] == '@' {
			total++
		}
	}
	return total < 4
}
