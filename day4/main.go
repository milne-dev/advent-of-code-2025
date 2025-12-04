package main

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

func main() {
	s, err := utils.ReadFileText("input.txt")
	if err != nil {
		panic(err)
	}

	grid := strings.Fields(s)
	fmt.Println(PartOne(grid))
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

var dirs = [][]int{{1, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}}

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
