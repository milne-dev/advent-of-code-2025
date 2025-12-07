package main

import (
	"aoc2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fields := strings.Fields(utils.ReadInput())
	fmt.Println(PartOne(fields))
}

func PartOne(fields []string) int {
	var ans int

	nums := make([]int, 0, len(fields))
	cmds := []byte{}

	for _, field := range fields {
		if field[0] == '*' || field[0] == '+' {
			cmds = append(cmds, field[0])
			continue
		}

		if n, err := strconv.Atoi(field); err != nil {
			panic(err)
		} else {
			nums = append(nums, n)
		}
	}

	rowSize := len(fields) / len(cmds)
	numsGrid := make([][]int, 0)
	for i := 0; i < len(nums); i += rowSize {
		numsGrid = append(numsGrid, nums[i:i+rowSize])
	}

	for i, cmd := range cmds {
		var x int
		if cmd == '*' {
			x = numsGrid[0][i]
			for j := 1; j < len(numsGrid); j++ {
				x *= numsGrid[j][i]
			}
		} else {
			for j := range numsGrid {
				x += numsGrid[j][i]
			}
		}
		ans += x
	}

	return ans
}

func PartTwo(fields []string) int {
	var ans int
	return ans
}
